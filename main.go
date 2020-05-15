package main

import (
	"github.com/crmartinez239/bqapi/controllers"
	"github.com/crmartinez239/bqapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/orders", controllers.FindOrders)
	r.POST("/orders", controllers.CreateOrder)

	r.GET("/orders/:id", controllers.FindOrder)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.DELETE("orders/:id", controllers.DeleteOrder)

	r.Run()
}
