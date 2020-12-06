package main

import (
	"github.com/gin-gonic/gin"
	  
	"github.com/tafaquh/mini-e-wallet/database"
	"github.com/tafaquh/mini-e-wallet/routes/auth"
	"github.com/tafaquh/mini-e-wallet/routes/common"
) 

func main() {
	err := database.ConnectMySQL()

	if err != nil {
		panic("DB connection failed!")
	}

	router := gin.Default()
	auth.Routes(router) //Added all auth routes
	common.Routes(router) //Added all auth routes
	router.Run()
}