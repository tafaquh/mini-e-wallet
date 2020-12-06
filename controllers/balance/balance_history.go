package balance

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tafaquh/mini-e-wallet/database"
	"github.com/tafaquh/mini-e-wallet/models/user_balance"
)

func GetBalanceHistory(c *gin.Context) { 
	user_id := c.Param("id")
	db := database.DBConn

	var ub user_balance.UserBalance
	if res := db.Preload("UserBalanceHistory").First(&ub, user_id); res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User balance not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get user balance",
		"data": ub,
	})
}