package model

import (
	"gorm.io/gorm"
)

type TempUser struct {
	gorm.Model
	FullName        string `gorm:"type:varchar(100);not null"`
	NoPhone         string `gorm:"type:varchar(20);not null"`
	EmailAddress    string `gorm:"type:varchar(100);not null"`
	Password        string `gorm:"type:varchar(100);not null"`
	ConfirmPassword string `gorm:"-"`
	VerificationCode string `gorm:"type:char(4);not null"`
	IsVerified      bool   `gorm:"type:boolean"`
}
