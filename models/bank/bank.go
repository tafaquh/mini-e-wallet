package bank

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	ID        	uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name  		string    `gorm:"size:255;not null" json:"name"`
	Location    string    `gorm:"size:100;not null" json:"location"`
}

