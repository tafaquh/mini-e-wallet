package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/tafaquh/mini-e-wallet/controllers/auth"
)

func Routes(route *gin.Engine){
	endpoint := route.Group("/api/v1/auth")
	{
		endpoint.GET("/logout", auth.Logout)
		endpoint.POST("/login", auth.Login)
	}
}