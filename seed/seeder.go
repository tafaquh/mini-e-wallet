package seed

import (
	"log"

	"gorm.io/gorm"
	"github.com/tafaquh/mini-e-wallet/models/user"
	"github.com/tafaquh/mini-e-wallet/models/bank"
	"github.com/tafaquh/mini-e-wallet/models/user_bank"
	"github.com/tafaquh/mini-e-wallet/models/user_balance"
	"github.com/tafaquh/mini-e-wallet/models/user_balance_history"
	"github.com/tafaquh/mini-e-wallet/models/bank_balance"
	"github.com/tafaquh/mini-e-wallet/models/bank_balance_history"
)

type Migrator interface {
	HasTable(dst interface{}) bool
}

var users = []user.User{
	user.User{
		Username: "Tafaquh",
		Email:    "tafaquh@gmail.com",
		Password: "tafaquh",
	},
	user.User{
		Username: "Fiddin",
		Email:    "fiddin@gmail.com",
		Password: "fiddin",
	},
	user.User{
		Username: "Orang",
		Email:    "orang@gmail.com",
		Password: "orang",
	},
}

var banks = []bank.Bank{
	bank.Bank{
		Name:   	"Bank Jago",
		Location: 	"Jalan Anak Jaksel, Jaksel",
	},
	bank.Bank{
		Name:  		"Bank Toktok",
		Location: 	"Jalan Tok tok tok, Jakpus",
	},
	bank.Bank{
		Name:  		"Bank Maju",
		Location: 	"Jalan Maju, Jakpus",
	},
}

var user_banks = []user_bank.UserBank{
	user_bank.UserBank{
		UserId: 1,
		BankId: 1,
		Status: "active",
	},
	user_bank.UserBank{
		UserId: 2,
		BankId: 1,
		Status: "active",
	},
	user_bank.UserBank{
		UserId: 3,
		BankId: 2,
		Status: "blocked",
	},
}

func Load(db *gorm.DB) {
	
	if err := db.Migrator().HasTable(&user.User{}); err { //only check user table, cannot use DropTableIfExists and not effective check all tables one by one
		return
	}

	err := db.Debug().AutoMigrate(&user.User{}, &bank.Bank{}, &user_bank.UserBank{}, &user_balance.UserBalance{}, &user_balance_history.UserBalanceHistory{}, &bank_balance.BankBalance{}, &bank_balance_history.BankBalanceHistory{})
	
	if err != nil {
		panic("cannot migrate tables!")
	}

	for i, _ := range users {
		hashedPassword, err := user.Hash(users[i].Password)
		users[i].Password = string(hashedPassword)
		if err != nil {
			log.Fatalf("cannot hash user password: %v", err)
		}

		err = db.Debug().Model(&user.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		err = db.Debug().Model(&bank.Bank{}).Create(&banks[i]).Error
		if err != nil {
			log.Fatalf("cannot seed banks table: %v", err)
		}

		err = db.Debug().Model(&user_bank.UserBank{}).Create(&user_banks[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user_banks table: %v", err)
		}
	}
}