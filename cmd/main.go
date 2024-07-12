package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/config"
	"api-gateway/logs"
	"api-gateway/service"
	"log"
)

func main() {
	cfg := config.Load()
	logs.InitLogger()
	logs.Logger.Info("Starting the application...")

	srv := service.New()
	h := handler.NewHandler(srv.Auth, srv.Reservation, srv.Payment, logs.Logger)

	r := api.NewRouter(h)

	logs.Logger.Info("Server is running", "PORT", cfg.HTTP_PORT)
	log.Fatal(r.Run(cfg.HTTP_PORT))
}
