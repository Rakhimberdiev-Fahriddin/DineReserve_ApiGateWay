package handler

import (
	pb "api-gateway/generated/payment_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePaymentHandler handles the creation of a new payment.
// @Summary Create Payment
// @Description Create a new payment record
// @Tags Payment
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Payment body payment_service.CreatePaymentRequest true "Create Payment"
// @Success 200 {object} payment_service.CreatePaymentResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/payments/ [post]
func (h *Handler) CreatePaymentHandler(ctx *gin.Context) {
    payment := pb.CreatePaymentRequest{}

    if err := ctx.ShouldBindJSON(&payment); err != nil {
        h.Logger.Error("Error binding JSON:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    resp, err := h.Payment.CreatePayment(ctx, &payment)
    if err != nil {
        h.Logger.Error("Error creating payment:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp.Payment)
}

// GetPaymentHandler retrieves the payment 
// @Summary Get payment 
// @Description Get a payment record
// @Tags Payment
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param payment-id path string true "payment-id"
// @Success 200 {object} payment_service.GetPaymentResponse
// @Failure 400 {object} string "Bad Request"
// @Router /api/payments/{payment-id} [get]
func (h *Handler) GetPaymentHandler(ctx *gin.Context) {
    h.Logger.Info("Handling GetPayment request")

    id := ctx.Param("payment-id")

    resp, err := h.Payment.GetPayment(ctx, &pb.GetPaymentRequest{Id: id})
    if err != nil {
        h.Logger.Error("Error getting payment: %v", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp.Payment)
}


// UpdatePaymentHandler updates the payment record
// @Summary Update Payment
// @Description Update the payment of record 
// @Tags Payment
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param payment-id path string true "Payment Id"
// @Param profile body payment_service.UpdatePaymentRequest true "Payment"
// @Success 200 {object} payment_service.UpdatePaymentResponse
// @Failure 400 {object} string "Bad Request"
// @Router /api/payments/{payment-id} [put]
func (h *Handler) UpdatePaymentHandler(ctx *gin.Context) {
    h.Logger.Info("Handling UpdatePayment request")

    id := ctx.Param("payment-id")

    payment := pb.UpdatePaymentRequest{}

    if err := ctx.ShouldBindJSON(&payment); err != nil {
        h.Logger.Error("Error binding JSON:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }
    payment.Id = id

    resp, err := h.Payment.UpdatePayment(ctx, &payment)
    if err != nil {
        h.Logger.Error("Error updating payment:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp.Payment)
}