package handlers

import (
	"net/http"

	"github.com/Alexander272/astro-atlas/internal/planet/models"
	"github.com/Alexander272/astro-atlas/internal/service"
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
	system := api.Group("/system")
	{
		system.GET("/", h.getSystemList)
		system.POST("/", h.createSystem)
		system.GET("/:id", h.getSystemById)
		system.PATCH("/:id", h.updateSystem)
		system.DELETE("/:id", h.deleteSystem)
	}

	planet := api.GET("/planet")
	{
		planet.GET("/", h.getPlanetList)
		planet.POST("/", h.createPlanet)
		planet.GET("/:id", h.getPlanetById)
		planet.PATCH("/:id", h.updatePlanet)
		planet.DELETE("/:id", h.deletePlanet)
	}
}

// @Summary Get System List
// @Tags systems
// @Description Получение списка планетных систем
// @ID getSystemList
// @Accept json
// @Produce json
// @Success 200 {array} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /system/ [get]
func (h *Handler) getSystemList(c *gin.Context) {
	systems, err := h.service.System.GetList(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: systems})
}

// @Summary Create System
// @Security ApiKeyAuth
// @Tags systems
// @Description Создание планетной системы
// @ID createSystem
// @Accept json
// @Produce json
// @Param system body models.CreateSystemDTO true "system info"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /system/ [post]
func (h *Handler) createSystem(c *gin.Context) {
	var dto models.CreateSystemDTO
	if err := c.BindJSON(&dto); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.service.System.Create(c, dto)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, idResponse{Id: id})
}

// @Summary Get System By Id
// @Tags systems
// @Description Получение планетной системы
// @ID getSystemById
// @Accept json
// @Produce json
// @Param id path string true "system id"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /system/{id} [get]
func (h *Handler) getSystemById(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	system, err := h.service.System.GetById(c, c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: system})
}

// @Summary Update System
// @Security ApiKeyAuth
// @Tags systems
// @Description Обновление данных планетной системы
// @ID updateSystem
// @Accept json
// @Produce json
// @Param id path string true "system id"
// @Param system body models.UpdateSystemDTO true "system info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /system/{id} [patch]
func (h *Handler) updateSystem(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
	}

	var dto models.UpdateSystemDTO
	if err := c.BindJSON(&dto); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.System.Update(c, dto)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "System updated"})
}

// @Summary Delete System
// @Security ApiKeyAuth
// @Tags systems
// @Description Удаление планетной системы
// @ID deleteSystem
// @Accept json
// @Produce json
// @Param id path string true "system id"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /system/{id} [delete]
func (h *Handler) deleteSystem(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	err := h.service.System.Delete(c, c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "System deleted"})
}

// @Summary Get Planet List
// @Tags planets
// @Description Получение списка планет системы
// @ID getPlanetList
// @Accept json
// @Produce json
// @Param system query string true "system id"
// @Success 200 {array} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /planet [get]
func (h *Handler) getPlanetList(c *gin.Context) {
	if c.Query("system") == "" {
		newResponse(c, http.StatusBadRequest, "empty system param")
		return
	}

	planets, err := h.service.Planet.GetList(c, c.Query("system"))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: planets})
}

// @Summary Create Planet
// @Security ApiKeyAuth
// @Tags planets
// @Description Создание планеты
// @ID createPlanet
// @Accept json
// @Produce json
// @Param planet body models.CreatePlanetDTO true "planet info"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /planet/ [post]
func (h *Handler) createPlanet(c *gin.Context) {
	var dto models.CreatePlanetDTO
	if err := c.BindJSON(&dto); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.service.Planet.Create(c, dto)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, idResponse{Id: id})
}

// @Summary Get Planet By Id
// @Tags planets
// @Description Получение планеты
// @ID getPlanetById
// @Accept json
// @Produce json
// @Param id path string true "planet id"
// @Success 200 {object} dataResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /planet/{id} [get]
func (h *Handler) getPlanetById(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	planet, err := h.service.Planet.GetById(c, c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: planet})
}

// @Summary updatePlanet
// @Security ApiKeyAuth
// @Tags planets
// @Description Обновление данных планеты
// @ID updatePlanet
// @Accept json
// @Produce json
// @Param id path string true "planet id"
// @Param planet body models.UpdatePlanetDTO true "planet info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /planet/{id} [patch]
func (h *Handler) updatePlanet(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	var dto models.UpdatePlanetDTO
	if err := c.BindJSON(&dto); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.Planet.Update(c, dto)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "Planet updated"})
}

// @Summary Delete Planet
// @Security ApiKeyAuth
// @Tags planets
// @Description Удаление планеты
// @ID deletePlanet
// @Accept json
// @Produce json
// @Param id path string true "planet id"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /planet/{id} [delete]
func (h *Handler) deletePlanet(c *gin.Context) {
	if c.Param("id") == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	err := h.service.Planet.Delete(c, c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "Planet deleted"})
}
