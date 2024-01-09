package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
)

type ProfileRepository struct{}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (r *ProfileRepository) CreateProfile(profile *model.Profile) error {
    db, err := database.Connect()
    if err != nil {
        return err
    }
    result := db.Create(profile)
    return result.Error
}

func (r *ProfileRepository) GetAllProfiles() ([]model.Profile, error) {
    db, err := database.Connect()
    if err != nil {
        return nil, err
    }
    var profiles []model.Profile
    result := db.Preload("User").Find(&profiles)
    return profiles, result.Error
}

func (r *ProfileRepository) GetProfile(id uint) (*model.Profile, error) {
    db, err := database.Connect()
    if err != nil {
        return nil, err
    }
    var profile model.Profile
    result := db.Preload("User").First(&profile, id)
    return &profile, result.Error
}

func (r *ProfileRepository) UpdateProfile(id uint, updatedProfile *model.Profile) error {
    db, err := database.Connect()
    if err != nil {
        return err
    }
    var profile model.Profile
    result := db.First(&profile, id)
    if result.Error != nil {
        return result.Error
    }
    result = db.Model(&profile).Updates(updatedProfile)
    return result.Error
}

func (r *ProfileRepository) DeleteProfile(id uint) error {
    db, err := database.Connect()
    if err != nil {
        return err
    }
    var profile model.Profile
    result := db.First(&profile, id)
    if result.Error != nil {
        return result.Error
    }
    result = db.Delete(&profile)
    return result.Error
}