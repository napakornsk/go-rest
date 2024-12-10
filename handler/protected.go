package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PortfolioHandler) ProtectedEndpointHandler(c *gin.Context) {
	userID := c.MustGet("user").(string)
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a protected endpoint",
		"userID":  userID,
	})
}
