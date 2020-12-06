package user_bank

import (
	"gorm.io/gorm"

	"github.com/tafaquh/mini-e-wallet/models/user"
	"github.com/tafaquh/mini-e-wallet/models/bank"
)

type UserBank struct {
	gorm.Model
	ID      uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	UserId	uint32    	`gorm:"not null" json:"user_id"`
	User 	user.User	`gorm:"foreignkey:UserId"`
	BankId	uint32    	`gorm:"not null" json:"bank_id"`
	Bank 	bank.Bank	`gorm:"foreignkey:BankId"`
	Status	string    	`gorm:"size:20;not null" json:"status"`
}

