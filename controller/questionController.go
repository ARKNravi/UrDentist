package controller

import (
	"net/http"
	"strconv"

	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var question model.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profileID, err := strconv.Atoi(c.Param("profileID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}
	question.ProfileID = uint(profileID)
	repo := repository.NewQuestionRepository()
	err = repo.CreateQuestion(&question)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Question created successfully"})
}

func GetAllQuestions(c *gin.Context) {
	repo := repository.NewQuestionRepository()
	questions, err := repo.GetAllQuestions()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, questions)
}

func GetQuestionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	repo := repository.NewQuestionRepository()
	question, err := repo.GetQuestionByID(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, question)
}