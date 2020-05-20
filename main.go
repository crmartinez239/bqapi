package main

import (
	"fmt"

	"github.com/crmartinez239/bqapi/controllers"
	"github.com/crmartinez239/bqapi/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	secretKey = "ACCESS_SECRET"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error in config: %s \n", err))
	}

	if !viper.IsSet(secretKey) {
		panic(fmt.Errorf("Fatal error %s not set\n", secretKey))
	}
}

func main() {
	r := gin.Default()
	ordersDB := models.SetupOrderModel()
	usersDB := models.SetupUserModel()

	r.Use(func(c *gin.Context) {
		c.Set("orders", ordersDB)
		c.Set("users", usersDB)
		c.Set("secret", viper.GetString(secretKey))
		c.Next()
	})

	r.GET("/orders", controllers.FindOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.FindOrder)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.DELETE("orders/:id", controllers.DeleteOrder)

	r.POST("/login", controllers.Login)
	r.Run()
}
