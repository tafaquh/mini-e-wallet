package common

import (
	"github.com/gin-gonic/gin"
	
	"github.com/tafaquh/mini-e-wallet/services/middleware"
	"github.com/tafaquh/mini-e-wallet/controllers/balance"
)

func Routes(route *gin.Engine){
	endpoint := route.Group("/api/v1/user/:id/balance")
	endpoint.Use(gin.Logger())
	endpoint.Use(gin.Recovery())
	endpoint.Use(middleware.TokenAuthMiddleware())
	{
		endpoint.GET("/balance-history", balance.GetBalanceHistory)
		endpoint.POST("/topup", balance.TopUp)
		endpoint.POST("/transfer", balance.Transfer)
	}
}