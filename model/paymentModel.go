package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Amount       float32 `gorm:"not null"`
	Status       bool    `gorm:"not null"`
	Method       string  `gorm:"type:varchar(100)"` 
	Photo        string  `gorm:"type:varchar(1024)"` 
	AppointmentID uint
}