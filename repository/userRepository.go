package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
	"gorm.io/gorm"
)

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
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

func (r *UserRepository) FindOrCreateUserByEmail(user *model.User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	if err := db.Where("email_address = ?", user.EmailAddress).First(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindUserByID(id string) (*model.User, error) {
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

func (r *UserRepository) UpdateUser(user *model.User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateTempUser(user *model.TempUser) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindTempUserByCode(code string) (*model.TempUser, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var user model.TempUser
	if err := db.Where("verification_code = ?", code).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no user found with verification code %s", code)
		}
		return nil, err
	}

	return &user, nil
}


func (r *UserRepository) MoveUserToDB(tempUser *model.TempUser) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	user := model.User{
		FullName:        tempUser.FullName,
		NoPhone:         tempUser.NoPhone,
		EmailAddress:    tempUser.EmailAddress,
		Password:        tempUser.Password,
		VerificationCode: tempUser.VerificationCode,
		IsVerified:      tempUser.IsVerified,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	if err := db.Delete(&tempUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUnverifiedUsers(tenMinutesAgo time.Time) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	err = db.Where("is_verified = ? AND created_at < ?", false, tenMinutesAgo).Delete(&model.TempUser{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindTempUserByEmail(email string) (*model.TempUser, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var user model.TempUser
	if err := db.Where("email_address = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateTempUser(user *model.TempUser) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
