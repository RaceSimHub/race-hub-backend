package middleware

import (
	"github.com/RaceSimHub/race-hub-backend/internal/model"
	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtenha o token do contexto, que foi configurado no JWTMiddleware
		claims := RetrieveJwtClaims(c)
		if claims == nil {
			return
		}

		// Verifique se o usuário possui a role de "admin"
		role, ok := claims["role"].(string)
		if !ok || role != string(model.UserRoleAdmin) {
			// Retorna erro no formato especificado
			response.Response{}.NewNotification(response.NotificationTypeError, "Acesso negado, você não tem permissão para acessar esta rota").WithRedirect("/").Show(c)
			c.Abort()
			return
		}

		// Se for admin, continue com a requisição
		c.Next()
	}
}

func RetrieveJwtClaims(c *gin.Context) jwt.MapClaims {
	token, exists := c.Get("token")
	if !exists {
		// Retorna erro no formato especificado
		response.Response{}.NewNotification(response.NotificationTypeError, "Token não encontrado").
			WithRedirect("/").
			Show(c)
		c.Abort()
		return nil
	}

	// Faz a asserção de tipo para o token
	parsedToken, ok := token.(*jwt.Token)
	if !ok {
		// Retorna erro no formato especificado
		response.Response{}.NewNotification(response.NotificationTypeError, "Token inválido").WithRedirect("/").Show(c)
		c.Abort()
		return nil
	}

	// Obtenha as claims do token (payload)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		// Retorna erro no formato especificado
		response.Response{}.NewNotification(response.NotificationTypeError, "Não foi possível extrair claims do token").WithRedirect("/").Show(c)
		c.Abort()
		return nil
	}

	return claims
}
