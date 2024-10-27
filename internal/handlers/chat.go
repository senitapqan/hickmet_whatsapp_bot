package handlers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MessageComes(c *gin.Context) {
	from := c.PostForm("From")
	body := c.PostForm("Body")

	phone := get_phone(from)

	if err := h.serv.Response(phone, body); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func get_phone(from string) string {
	phone := strings.TrimPrefix(from, "whatsapp:")
	re := regexp.MustCompile(`[^+\d]`)
	cleanedNumber := re.ReplaceAllString(phone, "")

	return cleanedNumber
}
