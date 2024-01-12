package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine) {
	r.POST("/profile", middleware.AuthMiddleware(), controller.CreateProfile)
	r.GET("/profile", middleware.AuthMiddleware(), controller.GetAllProfiles)
	r.GET("/profile/:profileID", middleware.AuthMiddleware(), controller.GetProfile)
	r.PUT("/profile/:profileID", middleware.AuthMiddleware(), controller.UpdateProfile)
	r.DELETE("/profile/:profileID", middleware.AuthMiddleware(), controller.DeleteProfile)
}
