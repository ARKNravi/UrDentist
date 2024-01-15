package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID          uint   `gorm:"not null"`
	NamaLengkap     string `gorm:"type:varchar(100);not null"`
	TempatLahir     string `gorm:"type:varchar(100);not null"`
	TanggalLahir    *time.Time `gorm:"type:date;not null"`
	Alamat          string `gorm:"type:varchar(255);not null"`
	Alergi          string `gorm:"type:varchar(255)"`
	User            User   `gorm:"foreignKey:UserID"`
	Ratings         []Rating `gorm:"foreignKey:ProfileID"`
	Questions       []Question `gorm:"foreignKey:ProfileID"`
	Appointment       []Appointment `gorm:"foreignKey:ProfileID"`
}

func (p *Profile) UnmarshalJSON(b []byte) error {
	type Alias Profile
	aux := &struct {
		TanggalLahir string `json:"tanggal_lahir"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	if aux.TanggalLahir != "" {
		t, err := time.Parse("2006-01-02", aux.TanggalLahir)
		if err != nil {
			return err
		}
		p.TanggalLahir = &t
	}
	return nil
}


