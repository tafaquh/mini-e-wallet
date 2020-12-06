package user_balance

import (
	"gorm.io/gorm"

	"github.com/tafaquh/mini-e-wallet/models/user"
)

type UserBalance struct {
	gorm.Model
	ID 				uint32  	`gorm:"primary_key;auto_increment" json:"id"`
	UserId  		uint32 		`gorm:"not null;" json:"user_id"`
	User 			user.User	`gorm:"foreignkey:UserId"`
	Balance 		int    		`gorm:"size:50;not null;" json:"balance"`
	BalanceAchieve	int    		`gorm:"size:50;not null;" json:"balance_achieve"`
}

