package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/crmartinez239/bqapi/models"
)

func FindOrders(c *gin.Context) {
	db := c.MustGet("orders").(*gorm.DB)

	var orders []models.OrderModel
	db.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func CreateOrder(c *gin.Context) {
	db := c.MustGet("orders").(*gorm.DB)

	var input models.CreateOrderModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.OrderModel{
		Type: input.Type,
		Name: input.Name,
		When: time.Now().String()}

	db.Create(&order)

	c.JSON(http.StatusOK, gin.H{"order": order})

}

func FindOrder(c *gin.Context) {
	db := c.MustGet("orders").(*gorm.DB)

	var order models.OrderModel

	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func UpdateOrder(c *gin.Context) {
	db := c.MustGet("orders").(*gorm.DB)

	var order models.OrderModel

	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found!"})
		return
	}

	var input models.UpdateOrderModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&order).Updates(input)

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func DeleteOrder(c *gin.Context) {
	db := c.MustGet("orders").(*gorm.DB)

	var order models.OrderModel
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete((&order))

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
