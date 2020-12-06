package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist"
	
	"github.com/tafaquh/mini-e-wallet/services/jwt"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		} 
		c.Next()
	}
}

// Authorize determines if current subject has been authorized to take an action on an object.
func Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "user hasn't logged in yet"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok := enforcer.Enforce(sub, obj, act)
	return ok, nil
}