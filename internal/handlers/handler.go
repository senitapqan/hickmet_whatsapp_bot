package handlers

import (
	"hickmet_whatapp_bot/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv service.Service
}

func NewHandler(serv service.Service) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/wa/bot")
	{
		api.GET("/get_user", h.GetUserByIIN)
	}
	return router
}
