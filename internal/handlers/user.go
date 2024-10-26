package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserByIIN(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "I got your first message",
	})
}
