package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Alexander272/astro-atlas/internal/service"
	"github.com/Alexander272/astro-atlas/internal/user/models"
	userService "github.com/Alexander272/astro-atlas/internal/user/service"
	"github.com/Alexander272/astro-atlas/pkg/apperror"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Services
}

func NewHandler(service *service.Services) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-out", h.signOut)
		auth.POST("/refresh", h.refresh)
	}

	// todo нужно добавить middleware
	users := api.Group("/users")
	{
		users.GET("/", h.getAllUsers)
		users.POST("/", h.createUser)
		users.GET("/:id", h.getUserById)
		users.PATCH("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
	}
}

// @Summary Sign In
// @Tags auth
// @Description вход в систему
// @ID signIn
// @Accept json
// @Produce json
// @Param signIn body models.SignInUserDTO true "credentials"
// @Success 200 {object} dataResponse{data=models.Token}
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var signIn models.SignInUserDTO
	if err := c.BindJSON(&signIn); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid credentials")
		return
	}

	ua := c.GetHeader("sec-ch-ua") + " " + c.GetHeader("sec-ch-ua-platform") + " " + c.GetHeader("User-Agent")
	ip := c.ClientIP()
	token, cookie, err := h.service.Auth.SignIn(c, signIn, ua, ip)
	if err != nil {
		if errors.Is(err, fmt.Errorf("invalid credentials")) || errors.Is(err, apperror.ErrNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	c.JSON(http.StatusOK, dataResponse{Data: token})
}

// @Summary Sign Out
// @Tags auth
// @Description выход из системы
// @ID signOut
// @Accept json
// @Produce json
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/sign-out [post]
func (h *Handler) signOut(c *gin.Context) {
	token, err := c.Cookie(userService.CookieName)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "cookie not found")
		return
	}

	cookie, err := h.service.Auth.SignOut(c, token)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	c.JSON(http.StatusOK, response{Message: "Successful sign out"})
}

// @Summary Refresh
// @Tags auth
// @Description обновление токенов доступа
// @Id refresh
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=models.Token}
// @Failure 400,403,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	ua := c.GetHeader("sec-ch-ua") + " " + c.GetHeader("sec-ch-ua-platform") + " " + c.GetHeader("User-Agent")
	ip := c.ClientIP()

	token, err := c.Cookie(userService.CookieName)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "cookie not found")
		return
	}

	netToken, cookie, err := h.service.Auth.Refresh(c, token, ua, ip)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	c.JSON(http.StatusOK, dataResponse{Data: netToken})
}

// @Summary Create User
// @Security ApiKeyAuth
// @Tags users
// @Description создание пользователя
// @ID createUser
// @Accept json
// @Produce json
// @Param user body models.CreateUserDTO true "user info"
// @Success 201 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var dto models.CreateUserDTO
	if err := c.BindJSON(&dto); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.service.User.Create(c, dto)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Header("Location", fmt.Sprintf("/api/users/%s", id))
	c.JSON(http.StatusCreated, idResponse{Id: id})
}

// @Summary Get All Users
// @Security ApiKeyAuth
// @Tags users
// @Description получение списка всех пользователей
// @ID getAllUsers
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]models.User}
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.User.GetAll(c)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dataResponse{Data: users})
}

// @Summary Get User By Id
// @Security ApiKeyAuth
// @Tags users
// @Description получение данных пользователя
// @ID getUserById
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dataResponse{data=models.User}
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	user, err := h.service.User.GetById(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: user})
}

// @Summary Update User
// @Security ApiKeyAuth
// @Tags users
// @Description обновление пользователя
// @ID updateUser
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param user body models.UpdateUserDTO true "user info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	var dto models.UpdateUserDTO
	if err := c.BindJSON(&dto); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	dto.Id = c.Param("id")

	err := h.service.User.Update(c, dto)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "User updated"})
}

// @Summary Delete User
// @Security ApiKeyAuth
// @Tags users
// @Description удаление пользователя
// @ID deleteUser
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 204 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	err := h.service.User.Delete(c, c.Param("id"))
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, response{Message: "User deleted"})
}
