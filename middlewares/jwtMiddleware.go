package middlewares

import (
	"net/http"
	"prog-web/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := services.ObterServicoAutenticacaoJWT().ValidarToken(tokenString)

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		} else if token != nil && !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
