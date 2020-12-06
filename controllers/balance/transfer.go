package balance

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tafaquh/mini-e-wallet/dto"
	"github.com/tafaquh/mini-e-wallet/database"
	"github.com/tafaquh/mini-e-wallet/models/user_balance"
)

func Transfer(c *gin.Context) { 
	user_id := c.Param("id")
	db := database.DBConn

	var transfer dto.Transfer
	err := c.BindJSON(&transfer)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Your transfer data are invalid!"})
		return
	}

	var ub user_balance.UserBalance
	if res := db.First(&ub, user_id); res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User balance not found."})
		return
	}

	if ub.Balance < transfer.Amount {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not enough user balance."})
		return
	}

	//update user balance history
	if err := SaveUserBalanceHistory(dto.CardType(1), int(ub.ID), ub.Balance, ub.Balance-transfer.Amount); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error})
		return
	}
	ub.Balance -= transfer.Amount
	ub.BalanceAchieve = -transfer.Amount

	db.Save(&ub) //update user balance

	//update user target balance history
	var ubt user_balance.UserBalance
	if res := db.First(&ubt, transfer.UserTargetId); res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User balance not found."})
		return
	}
	if err := SaveUserBalanceHistory(dto.CardType(1), int(ubt.ID), ubt.Balance, ubt.Balance+transfer.Amount); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error})
		return
	}
	ubt.Balance += transfer.Amount
	ubt.BalanceAchieve = transfer.Amount

	db.Save(&ubt) //update user target balance

	c.JSON(http.StatusOK, gin.H{
		"message": "transfer success",
		"your_balance": ub.Balance,
	})
}