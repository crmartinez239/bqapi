package main

import (
	"net/http"

	"github.com/crmartinez239/bqapi/controllers"
	"github.com/crmartinez239/bqapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(204)
		//	return
		//}
		c.Set("db", db)
		c.Next()
	})

	r.GET("/orders", controllers.FindOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.DELETE("/orders", controllers.DeleteAllOrders)

	r.GET("/orders/:id", controllers.FindOrder)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.DELETE("orders/:id", controllers.DeleteOrder)

	http.ListenAndServe(":9898", r)
}
