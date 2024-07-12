package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description This is a sample server for API Gateway.
// @host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func NewRouter(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.LoggerMiddleware())

	r := router.Group("/api")

	auth := r.Group("/auth")
	{
		auth.DELETE("/logout/:user-id", handle.LogoutUserHandler)
		auth.GET("/profile", handle.GetUserProfileHandler)
		auth.PUT("/profile/:user-id", handle.UpdateMenuItemHandler)
	}

	restaurant := r.Group("/restaurant")
	{
		restaurant.POST("/",handle.CreateRestaurantHandler)
		restaurant.GET("/",handle.ListRestaurantsHandler)
		restaurant.GET("/:restaurant-id",handle.GetRestaurantHandler)
		restaurant.PUT("/:restaurant-id",handle.UpdateRestaurantHandler)
		restaurant.DELETE("/:restaurant_id",handle.DeleteRestaurantHandler)
		
	}

	menu := r.Group("/menu")
	{
		menu.POST("/",handle.CreateMenuItemHandler)
		menu.GET("/",handle.ListMenuItemsHandler)
		menu.GET("/:menu-id",handle.GetMenuItemHandler)
		menu.PUT("/:menu-id",handle.UpdateMenuItemHandler)
		menu.DELETE("/:menu-id",handle.DeleteMenuItemHandler)
	}

	payment := r.Group("/payments")
	{
		payment.POST("/", handle.CreatePaymentHandler)
		payment.GET("/:payment-id", handle.GetPaymentHandler)
		payment.PUT("/:payment-id", handle.UpdatePaymentHandler)
	}

	reservation := r.GET("/reservations")
	{
		reservation.POST("/",handle.CreateReservationHandler)
		reservation.GET("/",handle.ListReservationHandler)
		reservation.GET("/:reservation-id",handle.GetReservationHandler)
		reservation.PUT("/:reservation-id",handle.UpdateReservationHandler)
		reservation.DELETE("/:reservation-id",handle.DeleteReservationHandler)
		reservation.POST("/check",handle.CheckReservationHandler)
		reservation.POST("/order",handle.OrderMealsHandler)
	}

	return router
}
