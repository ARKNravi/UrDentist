package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
)

type DentistRepository struct {}

func NewDentistRepository() *DentistRepository {
	return &DentistRepository{}
}

func (r *DentistRepository) CreateDentist(dentist *model.Dentist) error {
    db, err := database.Connect()
    if err != nil {
        return err
    }
    result := db.Create(dentist)
    return result.Error
}

func (r *DentistRepository) GetAllDentists() ([]model.Dentist, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var dentists []model.Dentist
	result := db.Find(&dentists)
	return dentists, result.Error
}

func (r *DentistRepository) GetDentistByID(id uint) (*model.Dentist, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var dentist model.Dentist
	result := db.First(&dentist, id)
	return &dentist, result.Error
}
