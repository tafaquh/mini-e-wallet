package bank_balance_history

import (
	"gorm.io/gorm"

	"github.com/tafaquh/mini-e-wallet/models/bank_balance"
	"github.com/tafaquh/mini-e-wallet/dto"
)

type BankBalanceHistory struct {
	gorm.Model
	ID 				uint32  					`gorm:"primary_key;auto_increment" json:"id"`
	BankBalanceId  	uint32 						`gorm:"not null;" json:"bank_balance_id"`
	BalanceBefore	int    						`gorm:"size:50;not null;" json:"balance_before"`
	BalanceAfter	int    						`gorm:"size:50;not null;" json:"balance_after"`
	Activity		string    					`gorm:"size:50;not null;" json:"activity"`
	Type			dto.CardType    			`gorm:"not null;" json:"type" sql:"type:ENUM('credit', 'debit')"`
	Ip				string    					`gorm:"size:50;" json:"ip"`
	Location		string    					`gorm:"size:50;" json:"location"`
	UserAgent		string    					`gorm:"size:50;" json:"user_agent"`
	Author			string    					`gorm:"size:50;" json:"author"`
	BankBalance		bank_balance.BankBalance	`gorm:"foreignkey:BankBalanceId"`
}

