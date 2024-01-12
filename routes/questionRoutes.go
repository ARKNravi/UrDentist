package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/gin-gonic/gin"
)

func QuestionRoutes(r *gin.Engine) {
	r.POST("/profile/:profileID/question", middleware.AuthMiddleware(), controller.CreateQuestion)
	r.GET("/questions", middleware.AuthMiddleware(), controller.GetAllQuestions)
	r.GET("/question/:id", middleware.AuthMiddleware(), controller.GetQuestionByID)
	r.PUT("/dentist/:dentistID/question/:id/answer", middleware.AuthMiddleware(), controller.AnswerQuestion)
}