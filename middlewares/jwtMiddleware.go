package middlewares

import (
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// success, err := services.Login(c.GetHeader("usuario"), c.GetHeader("senha"))
		// if err != nil {
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		// }

		// if !success {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }
	}
}
