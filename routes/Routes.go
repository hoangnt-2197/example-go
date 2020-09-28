package routes

import (
	"example/auth"
	"example/controllers"
	"example/middlewares"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	e, err := casbin.NewEnforcer("auth_model.conf", "auth_policy.csv")
	fmt.Println("---------", err)

	grp1 := r.Group("user-api")

	{
		grp1.POST("login", controllers.Login)
	}

	grp1.Use(middlewares.JWTTokenFilter)
	grp1.Use(auth.NewAuthorizer(e))
	{
		grp1.GET("users", controllers.GetUsers)
		grp1.POST("user", controllers.CreateUser)
		grp1.GET("user/:id", controllers.GetUserByID)
		grp1.PUT("user/:id", controllers.UpdateUser)
		grp1.DELETE("user/:id", controllers.DeleteUser)
	}

	return r
}
