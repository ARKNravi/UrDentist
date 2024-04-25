package controller

import (
	"net/http"
	"strconv"

	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

func CreateDentist(c *gin.Context) {
    var dentist model.Dentist
    if err := c.ShouldBindJSON(&dentist); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    repo := repository.NewDentistRepository()
    err := repo.CreateDentist(&dentist)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Profile created successfully"})
}

func GetAllDentists(c *gin.Context) {
	repo := repository.NewDentistRepository()
	dentists, err := repo.GetAllDentists()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, dentists)
}

func GetDentistByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	repo, err := repository.NewAppointmentRepository()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	DentistRepository:= repository.NewDentistRepository()
	dentist, err := DentistRepository.GetDentistByID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	offlineConsultations, err := repo.GetOfflineConsultationsByDentistID(dentist.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	onlineConsultations, err := repo.GetOnlineConsultationsByDentistID(dentist.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	dentist.OfflineConsultations = offlineConsultations
	dentist.OnlineConsultations = onlineConsultations
	c.JSON(200, dentist)
}
