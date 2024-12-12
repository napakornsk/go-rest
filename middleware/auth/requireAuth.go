package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *AuthService) RequireAuth(c *gin.Context) {
	// Get the cookie off req
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer")

	claims, err := h.validateToken(&tokenString)
	if err != nil {
		log.Printf("Token is invalid: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	if h.isTokenExpire(claims) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	user, err := h.FindUser(claims["username"].(string))
	if err != nil {
		log.Printf("Invalid credential: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credential"})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()

}
