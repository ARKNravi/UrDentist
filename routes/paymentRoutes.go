// routes/routes.go
package routes

import (
	"log"

	"github.com/ARKNravi/HACKFEST-BE/controller"
	"github.com/ARKNravi/HACKFEST-BE/middleware"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	repo, err := repository.NewPaymentRepository()
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	controller := controller.NewPaymentController(repo)
	r.PUT("/paymentsdummy/:paymentID", middleware.AuthMiddleware(), controller.UpdatePaymentDummy)
	r.PUT("/payments/:paymentID", middleware.AuthMiddleware(), controller.UpdatePayment)
}
