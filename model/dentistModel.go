package model

import (
	"gorm.io/gorm"
)

type Dentist struct {
	gorm.Model
	Name                string `gorm:"type:varchar(100);not null"`
	Specialist          string `gorm:"type:varchar(100);not null"`
	AboutMe             string `gorm:"type:text;not null"`
	WorkYearExperience  int    `gorm:"not null"`
	PatientCount        int
	Picture             string `gorm:"type:varchar(255)"`
	OnlineConsultations []OnlineConsultation `gorm:"foreignKey:DentistID"`
	OfflineConsultations []OfflineConsultation `gorm:"foreignKey:DentistID"`
	Appointments        []Appointment `gorm:"foreignKey:DentistID"`
	Ratings             []Rating `gorm:"foreignKey:DentistID"`
	Questions           []Question `gorm:"foreignKey:DentistID"`
}

