package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/panjiasmoroart/goforum26/internal/configs"
	"github.com/panjiasmoroart/goforum26/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			// c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing token",
			})
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			// c.AbortWithError(http.StatusUnauthorized, err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		// jika kita mempunyai 2 middleware akan lanjut ke middleware selnjtunya
		// jika tidak di panggil makan akan berhenti pada middleware ini saja
		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
