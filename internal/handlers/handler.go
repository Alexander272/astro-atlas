package handlers

import (
	"fmt"
	"net/http"

	"github.com/Alexander272/astro-atlas/internal/config"
	planetHandlers "github.com/Alexander272/astro-atlas/internal/planet/handlers"
	"github.com/Alexander272/astro-atlas/internal/service"
	userHandlers "github.com/Alexander272/astro-atlas/internal/user/handlers"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag/example/basic/docs"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(cors.Config{
			AllowedOrigins: []string{conf.Http.Host},
			AllowedMethods: []string{"GET"},
			AllowedHeaders: []string{"Origin"},
			ExposedHeaders: []string{"Content-Length"},
			// AllowCredentials: true,
		}),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", conf.Http.Host, conf.Http.Port)
	if conf.Environment != "dev" {
		docs.SwaggerInfo.Host = conf.Http.Host
	}

	if conf.Environment != "prod" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Init router
	router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	planetHandler := planetHandlers.NewHandler(h.services)
	userHandler := userHandlers.NewHandler(h.services)

	api := router.Group("/api")
	{
		userHandler.Init(api)
		planetHandler.Init(api)
	}
}
