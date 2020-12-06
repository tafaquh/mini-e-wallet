package bank_balance

import (
	"gorm.io/gorm"

	"github.com/tafaquh/mini-e-wallet/models/bank"
)

type BankBalance struct {
	gorm.Model
	ID 				uint32  	`gorm:"primary_key;auto_increment" json:"id"`
	BankId  		uint32 		`gorm:"not null" json:"bank_id"`
	Balance 		int    		`gorm:"size:50;not null" json:"balance"`
	BalanceAchieve	int    		`gorm:"size:50;not null" json:"balance_achieve"`
	Code			string  	`gorm:"size:50" json:"code"`
	Enable			bool   		`gorm:"size:50" json:"enable"`
	Bank 			bank.Bank	`gorm:"foreignkey:BankId"`
}

