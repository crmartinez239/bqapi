package controllers

import (
	"net/http"

	"github.com/crmartinez239/bqapi/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Login(c *gin.Context) {
	var input models.UserModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid JSON")
		return
	}

	db := c.MustGet("users").(*gorm.DB)
	var local models.UserModel
	if err := db.Where("username = ?", input.Username).First(&local).Error; err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid login details")
		return
	}

	if input.Password != local.Password {
		c.JSON(http.StatusUnauthorized, "Invalid login details")
		return
	}

	// create auth token
	secret := c.MustGet("secret").(string)

	c.JSON(http.StatusOK, gin.H{"secret": secret})

}
