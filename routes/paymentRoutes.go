// routes/routes.go
package routes

import (
	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	r.PUT("/paymentsdummy/:paymentID", middleware.AuthMiddleware(), controller.UpdatePaymentDummy)
	r.PUT("/payments/:paymentID", middleware.AuthMiddleware(), controller.UpdatePayment)
}
