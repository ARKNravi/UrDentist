package controller

import (
	"net/http"
	"strconv"

	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

func CreateProfile(c *gin.Context) {
    var profile model.Profile
    if err := c.ShouldBindJSON(&profile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    userID := uint(c.MustGet("userID").(float64))
    profile.UserID = userID
    repo := repository.NewProfileRepository()
    err := repo.CreateProfile(&profile)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Profile created successfully"})
}

func GetAllProfiles(c *gin.Context) {
    repo := repository.NewProfileRepository()
    profiles, err := repo.GetAllProfiles()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, profiles)
}

func GetProfile(c *gin.Context) {
    id := c.Param("id")
    idUint, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid profile ID"})
        return
    }
    repo := repository.NewProfileRepository()
    profile, err := repo.GetProfile(uint(idUint))
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, profile)
}

func UpdateProfile(c *gin.Context) {
    id := c.Param("id")
    idUint, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid profile ID"})
        return
    }
    var updatedProfile model.Profile
    if err := c.ShouldBindJSON(&updatedProfile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    repo := repository.NewProfileRepository()
    err = repo.UpdateProfile(uint(idUint), &updatedProfile)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Profile updated successfully"})
}

func DeleteProfile(c *gin.Context) {
    id := c.Param("id")
    idUint, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid profile ID"})
        return
    }
    repo := repository.NewProfileRepository()
    err = repo.DeleteProfile(uint(idUint))
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Profile deleted successfully"})
}
