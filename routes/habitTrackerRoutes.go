package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	r.PUT("/profile/:profileID/task/:taskID/undo", middleware.AuthMiddleware(), controller.UndoTask)
    r.GET("/profile/:profileID/tasks", middleware.AuthMiddleware(), controller.GetTasksByDate)
    r.POST("/profile/:profileID/task/:taskID/complete", middleware.AuthMiddleware(), controller.CompleteTask)
	r.GET("/tasks", controller.GetAllTasks)
	r.GET("/profile/:profileID/completedTasks", middleware.AuthMiddleware(), controller.GetCompletedTasks)
}
