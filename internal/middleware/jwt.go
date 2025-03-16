package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RaceSimHub/race-hub-backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte(config.JwtSecret)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/js/") || strings.HasPrefix(path, "/css/") || path == "/favicon.ico" {
			c.Next()
			return
		}

		tokenStr, err := c.Cookie("jwt")
		if err != nil || tokenStr == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
