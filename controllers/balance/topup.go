package balance

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tafaquh/mini-e-wallet/dto"
	"github.com/tafaquh/mini-e-wallet/database"
	"github.com/tafaquh/mini-e-wallet/models/user_bank"
	"github.com/tafaquh/mini-e-wallet/models/user_balance"
	"github.com/tafaquh/mini-e-wallet/models/user_balance_history"
	"github.com/tafaquh/mini-e-wallet/models/bank_balance"
	"github.com/tafaquh/mini-e-wallet/models/bank_balance_history"
)

type CardType string

const (
	Credit CardType = "credit"
	Debit CardType = "debit"
)

func TopUp(c *gin.Context) { 
	user_id := c.Param("id")
	db := database.DBConn

	var ubank user_bank.UserBank
	if res := db.First(&ubank, user_id); res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User doesnt have any bank account."})
		return
	}

	if ubank.Status != "active" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User back account not active! Call your bank."})
		return
	}

	var amount dto.TopUp
	err := c.BindJSON(&amount)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Your balance data type are invalid!"})
		return 
	}

	var ub user_balance.UserBalance
	if res := db.First(&ub, user_id); res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User balance not found."})
		return
	}

	var bb bank_balance.BankBalance
	if res := db.First(&bb, ubank.BankId); res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Bank balance not found."})
		return
	}

	if err := SaveBankBalanceHistory(ubank.Type, int(ubank.BankId), bb.Balance, bb.Balance-amount.BalanceAchieve); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error})
	}
	bb.Balance -= amount.BalanceAchieve
	bb.BalanceAchieve = amount.BalanceAchieve

	if err := SaveUserBalanceHistory(ubank.Type, int(ub.ID), ub.Balance, ub.Balance+amount.BalanceAchieve); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error})
	}
	ub.Balance += amount.BalanceAchieve
	ub.BalanceAchieve = amount.BalanceAchieve

	db.Save(&ub)
	db.Save(&bb)

	c.JSON(http.StatusOK, gin.H{
		"message": "topup success",
		"balance": ub.Balance,
	})
}

func SaveUserBalanceHistory(ctype dto.CardType, user_balance_id, balance_before, balance_after int) error {
	db := database.DBConn
	
	var ubh user_balance_history.UserBalanceHistory
	ubh.UserBalanceId = uint32(user_balance_id)
	ubh.BalanceBefore = balance_before
	ubh.BalanceAfter = balance_after
	ubh.Activity = "topup"
	ubh.Type = ctype

	// Check if there's error when insert to database
	if err := db.Create(&ubh); err.Error != nil {
		return errors.New("Failed to store to database")
	}

	return nil
}

func SaveBankBalanceHistory(ctype dto.CardType, bank_balance_id, balance_before, balance_after int) error {
	db := database.DBConn

	if balance_after < 0 {
		return errors.New("Not enough balance!")
	}

	var bbh bank_balance_history.BankBalanceHistory
	bbh.BankBalanceId = uint32(bank_balance_id)
	bbh.BalanceBefore = balance_before
	bbh.BalanceAfter = balance_after
	bbh.Activity = "topup"
	bbh.Type = ctype

	// Check if there's error when insert to database
	if err := db.Create(&bbh); err.Error != nil {
		return errors.New("Failed to store to database")
	}

	return nil
}