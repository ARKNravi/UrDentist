package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	userController := controller.UserController{}
	route.POST("/register", userController.Register)
	route.POST("/login", userController.Login)
	route.GET("/auth/google/login", userController.GoogleLogin)
	route.GET("/auth/google/callback", userController.GoogleCallback)
	route.POST("/verify", userController.Verify)
	route.POST("/resend-verification", userController.ResendVerificationCode)
	route.POST("/forgot-password", userController.ForgotPassword)
	route.POST("/verify-token", userController.VerifyToken)
	route.POST("/reset-password", middleware.AuthMiddlewarePassword(), userController.ResetPassword)
}




