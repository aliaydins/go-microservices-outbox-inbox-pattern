package order

import (
	"github.com/aliaydins/oipattern/shared/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	h.iRoutes(router)
	return router
}

func (h *Handler) iRoutes(router *gin.Engine) {
	router.Use(middleware.CORS())
	routerGroup := router.Group("/")
	routerGroup.GET("/health", h.health)
	routerGroup.GET("/outbox", h.getList)
	routerGroup.POST("/order", h.createOrder)
}
