package middlewares

import (
	"example/auth"
	"example/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTTokenFilter(c *gin.Context) {
	err := auth.TokenValid(c.Request)
	if err != nil {
		message := utils.MessageError {
			Code: 401,
			Message: "unauthorized",
		}

		c.JSON(http.StatusUnauthorized,message)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}