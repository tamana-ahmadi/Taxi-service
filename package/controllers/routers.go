package controllers

import (
	"Taxi_service/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)
	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}
	usersG := router.Group("/users", checkUserAuthentication)
	{
		usersG.POST("", AddUsers)
		usersG.GET("", PrintUsers)
		usersG.GET("/:id", PrintUsersByID)
		usersG.PUT("/:id", EditUsers)
		usersG.PATCH("/:id", EditUsersPassword)
		usersG.DELETE("/:id", DeleteUsers)
	}

	routesG := router.Group("/routes", checkUserAuthentication)
	{
		routesG.POST("", AddRoute)
		routesG.GET("", GetAllRoutes)
		routesG.GET("/:id", GetAllRoutesByID)
		routesG.PUT("/:id", UpdateRouteByID)
		routesG.PATCH("/:id", ChecksRouteasResponse)
		routesG.DELETE("/:id", DeleteRouteByID)
	}

	taxicompsG := router.Group("/taxicomps", checkUserAuthentication)
	{
		taxicompsG.POST("", AddTaxicomp)
		taxicompsG.GET("", GetAllTaxiComp)
		taxicompsG.GET("/:id", GetAllTaxiCompByID)
		taxicompsG.PUT("/:id", UpdateTaxiCompByID)
		taxicompsG.DELETE("/:id", DeleteTaxiCompByID)
	}
	ordersG := router.Group("/orders", checkUserAuthentication)
	{
		ordersG.POST("", AddOrder)
		ordersG.GET("", OrdersReport)
		ordersG.GET("/:id", GetAllOrdersByID)
		ordersG.PUT("/:id", UpdateOrderByID)
		ordersG.PATCH("/:id", ChecksOrderasDone)
		ordersG.DELETE("/:id", DeleteOrderByID)
	}

	return router
}
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
