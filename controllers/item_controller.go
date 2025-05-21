package controllers

import (
	"aplikasi-daftar-belanja/config"
	"aplikasi-daftar-belanja/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"user not found in context"})
	}
	config.DB.Where("user_id = ?", userID).Find(&items)
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var item models.Item
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"user not found in context"})
	}
	if err := c.ShouldBindBodyWithJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.UserID = uint(userID.(float64))
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context){
	id := c.Param("id")
	var item models.Item
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"user not found in context"})
	}
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Find(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input struct {
		Name 				*string `json:"name"`
		Quantity 		*int 		`json:"quantity"`
		IsPurchased *bool 	`json:"is_purchased"`
		Note        *string `json:"note"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != nil {
		item.Name = *input.Name
	}
	if input.Quantity != nil {
		item.Quantity = *input.Quantity
	}
	if input.IsPurchased != nil {
		item.IsPurchased = *input.IsPurchased
		if !item.IsPurchased {
			item.PurchasedAt = nil
		} else {
			now := time.Now()
			item.PurchasedAt = &now
		}
	}
	if input.Note != nil {
		item.Note = *input.Note
	}

	config.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"user not found in context"})
	}
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Find(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message":"Item deleted"})
}