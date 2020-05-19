package main

import (
	"github.com/crmartinez239/bqapi/controllers"
	"github.com/crmartinez239/bqapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	ordersDB := models.SetupOrderModel()
	//usersDB := models.SetupUserModel()

	r.Use(func(c *gin.Context) {
		c.Set("orders", ordersDB)
		//c.Set("users", usersDB)
		c.Next()
	})

	r.GET("/orders", controllers.FindOrders)
	r.POST("/orders", controllers.CreateOrder)

	r.GET("/orders/:id", controllers.FindOrder)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.DELETE("orders/:id", controllers.DeleteOrder)

	r.Run()
}
