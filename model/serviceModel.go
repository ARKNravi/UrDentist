package model

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text"`
	Duration    int     `gorm:"not null"`
	OnlineConsultations []OnlineConsultation `gorm:"foreignKey:ServiceID"`
	OfflineConsultations []OfflineConsultation `gorm:"foreignKey:ServiceID"`
}