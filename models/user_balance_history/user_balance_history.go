package user_balance_history

import (
	"gorm.io/gorm"

	"github.com/tafaquh/mini-e-wallet/models/user_balance"
	"github.com/tafaquh/mini-e-wallet/dto"
)

type CardType string

const (
	Credit CardType = "credit"
	Debit CardType = "debit"
)

type UserBalanceHistory struct {
	gorm.Model
	ID 				uint32  					`gorm:"primary_key;auto_increment" json:"id"`
	UserBalanceId  	uint32 						`gorm:"not null;" json:"user_balance_id"`
	BalanceBefore	int    						`gorm:"size:50;not null;" json:"balance_before"`
	BalanceAfter	int    						`gorm:"size:50;not null;" json:"balance_after"`
	Activity		string    					`gorm:"size:50;not null;" json:"activity"`
	Type			dto.CardType    			`gorm:"not null;" json:"type" sql:"type:ENUM('credit', 'debit')"`
	Ip				string    					`gorm:"size:50;" json:"ip"`
	Location		string    					`gorm:"size:50;" json:"location"`
	UserAgent		string    					`gorm:"size:50;" json:"user_agent"`
	Author			string    					`gorm:"size:50;" json:"author"`
	UserBalance		user_balance.UserBalance	`gorm:"foreignkey:UserBalanceId"`
}

