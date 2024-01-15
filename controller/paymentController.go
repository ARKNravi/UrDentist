// controller/paymentController.go
package controller

import (
	"net/http"
	"strconv"

	"github.com/ARKNravi/HACKFEST-BE/model"
	"github.com/ARKNravi/HACKFEST-BE/repository"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	repo repository.PaymentRepository
}

func NewPaymentController(repo repository.PaymentRepository) *PaymentController {
	return &PaymentController{repo: repo}
}

func (c *PaymentController) UpdatePayment(ctx *gin.Context) {
	paymentID, err := strconv.Atoi(ctx.Param("paymentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var payment model.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.ID = uint(paymentID)

	if err := c.repo.Update(&payment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the payment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"payment": payment})
}
