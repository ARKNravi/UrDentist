package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName        string `gorm:"type:varchar(100);not null"`
	NoPhone         string `gorm:"type:varchar(20);not null;unique"`
	EmailAddress    string `gorm:"type:varchar(100);not null;unique"`
	Password        string `gorm:"type:varchar(100);not null"`
	ConfirmPassword string `gorm:"-"`
	VerificationCode string `gorm:"type:char(4);not null"`
	IsVerified      bool   `gorm:"type:boolean"`
	Profiles []Profile `gorm:"foreignKey:UserID"`
}