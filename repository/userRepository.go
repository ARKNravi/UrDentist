package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
)

func FindUserByEmail(email string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := db.Where("email_address = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *model.User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	
	return nil
}

func FindOrCreateUserByEmail(user *model.User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	if err := db.Where("email_address = ?", user.EmailAddress).FirstOrCreate(&user).Error; err != nil {
		return err
	}

	return nil
}

func FindUserByID(id string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var user model.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *model.User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}