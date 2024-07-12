package handler

import (
	user "api-gateway/generated/auth_service"
	payment "api-gateway/generated/payment_service"
	reservation "api-gateway/generated/reservation_service"
	"log/slog"
)

type Handler struct {
	Logger      *slog.Logger
	Auth        user.AuthServiceClient
	Reservation reservation.ReservationServiceClient
	Payment     payment.PaymentServiceClient
}

func NewHandler(authClient user.AuthServiceClient, reservationClient reservation.ReservationServiceClient, paymentClient payment.PaymentServiceClient, logger *slog.Logger) *Handler {
	return &Handler{
		Auth:        authClient,
		Reservation: reservationClient,
		Payment:     paymentClient,
		Logger:      logger,
	}
}
