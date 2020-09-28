package auth

import (
	"example/models"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BasicAuthorizer struct {
	enforcer * casbin.Enforcer
}


func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	authorizer := &BasicAuthorizer{enforcer: e}
	return func(context *gin.Context) {
		if !authorizer.CheckPermission(context.Request) {
			authorizer.RequirePermission(context)
		}
	}
}

func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	userId, err := ExtractTokenID(r)
	if err != nil {
		return false
	}
	var user models.User
	method := r.Method
	path := r.URL.Path
	err = models.GetUserByID(&user, fmt.Sprint(userId))
	if err != nil {
		return false
	}
	allowed, err := a.enforcer.Enforce(user.Role.Name, path, method)
	if err != nil {
		panic(err)
	}
	return allowed
}

func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}