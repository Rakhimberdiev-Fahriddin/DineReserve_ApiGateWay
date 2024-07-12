package handler

import (
	pb "api-gateway/generated/reservation_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateReservationHandler handles the creation of a new reservation.
// @Summary Create Reservation
// @Description Create a new reservation
// @Tags Reservation
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Reservation body reservation_service.CreateReservationRequest true "Create Reservation"
// @Success 200 {object} reservation_service.CreateReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations [post]
func (h *Handler) CreateReservationHandler(ctx *gin.Context) {
    h.Logger.Info("Handling CreateReservation request")

    reservation := pb.CreateReservationRequest{}
    if err := ctx.ShouldBindJSON(&reservation); err != nil {
        h.Logger.Error("Error binding JSON:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    resp, err := h.Reservation.CreateReservation(ctx, &reservation)
    if err != nil {
        h.Logger.Error("Error creating reservation:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp.Reservation)
}

// ListReservationHandler lists all reservations.
// @Summary List Reservations
// @Description List all reservations
// @Tags Reservation
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} reservation_service.ListReservationsResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations [get]
func (h *Handler) ListReservationHandler(ctx *gin.Context) {
    h.Logger.Info("Handling ListReservation request")

    var filter pb.ListReservationsRequest
    if err := ctx.ShouldBindQuery(&filter); err != nil {
        h.Logger.Error("Error binding query parameters:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    if filter.Limit == 0 {
        filter.Limit = 10
    }

    resp, err := h.Reservation.ListReservations(ctx, &filter)
    if err != nil {
        h.Logger.Error("Error listing reservations:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp)
}
// GetReservationHandler retrieves a specific reservation by ID.
// @Summary Get Reservation
// @Description Get a specific reservation by ID
// @Tags Reservation
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param reservation-id path string true "Reservation ID"
// @Success 200 {object} reservation_service.GetReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/{reservation-id} [get]
func (h *Handler) GetReservationHandler(ctx *gin.Context) {
    h.Logger.Info("Handling GetReservation request")

    id := ctx.Param("reservation-id")

    resp, err := h.Reservation.GetReservation(ctx, &pb.GetReservationRequest{Id: id})
    if err != nil {
        h.Logger.Error("Error getting reservation:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp.Reservation)
}

// UpdateReservationHandler updates an existing reservation.
// @Summary Update Reservation
// @Description Update an existing reservation
// @Tags Reservation
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param reservation-id path string true "Reservation ID"
// @Param Reservation body reservation_service.UpdateReservationRequest true "Update Reservation"
// @Success 200 {object} reservation_service.UpdateReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/{reservation-id} [put]
func (h *Handler) UpdateReservationHandler(ctx *gin.Context) {
    h.Logger.Info("Handling UpdateReservation request")

    id := ctx.Param("reservation-id")

    reservation := pb.UpdateReservationRequest{}
    if err := ctx.ShouldBindJSON(&reservation); err != nil {
        h.Logger.Error("Error binding JSON:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }
    reservation.Id = id

    resp, err := h.Reservation.UpdateReservation(ctx, &reservation)
    if err != nil {
        h.Logger.Error("Error updating reservation:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp.Reservation)
}


// DeleteReservationHandler deletes a specific reservation by ID.
// @Summary Delete Reservation
// @Description Delete a specific reservation by ID
// @Tags Reservation
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param reservation-id path string true "Reservation ID"
// @Success 200 {object} reservation_service.DeleteReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/{reservation-id} [delete]
func (h *Handler) DeleteReservationHandler(ctx *gin.Context) {
    h.Logger.Info("Handling DeleteReservation request")

    id := ctx.Param("reservation-id")

    reservation := pb.DeleteReservationRequest{}
    if err := ctx.ShouldBindJSON(&reservation); err != nil {
        h.Logger.Error("Error binding JSON:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }
    reservation.Id = id

    resp, err := h.Reservation.DeleteReservation(ctx, &reservation)
    if err != nil {
        h.Logger.Error("Error deleting reservation:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp.Message)
}


// CheckReservationHandler checks the status of a reservation.
// @Summary Check Reservation
// @Description Check the status of a reservation
// @Tags Reservation
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Check body reservation_service.CheckReservationRequest true "Check Reservation"
// @Success 200 {object} reservation_service.CheckReservationResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/check [post]
func (h *Handler) CheckReservationHandler(ctx *gin.Context) {
    h.Logger.Info("Handling CheckReservation request")

    checkReq := pb.CheckReservationRequest{}
    if err := ctx.ShouldBindJSON(&checkReq); err != nil {
        h.Logger.Error("Error binding JSON:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    resp, err := h.Reservation.CheckReservation(ctx, &checkReq)
    if err != nil {
        h.Logger.Error("Error checking reservation:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp)
}


// OrderMealsHandler handles ordering meals for a reservation.
// @Summary Order Meals
// @Description Order meals for a reservation
// @Tags Reservation
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Order body reservation_service.OrderMealsRequest true "Order Meals"
// @Success 200 {object} reservation_service.OrderMealsResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/reservations/order [post]
func (h *Handler) OrderMealsHandler(ctx *gin.Context) {
    h.Logger.Info("Handling OrderMeals request")

    orderReq := pb.OrderMealsRequest{}
    if err := ctx.ShouldBindJSON(&orderReq); err != nil {
        h.Logger.Error("Error binding JSON:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    resp, err := h.Reservation.OrderMeals(ctx, &orderReq)
    if err != nil {
        h.Logger.Error("Error ordering meals:", "error", err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp)
}
