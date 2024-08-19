package handlers

import (
	"gin-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetItems mendapatkan semua item dari database
func GetItems(c *gin.Context) {
    items, err := models.GetItems()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, items)
}

// GetItem mendapatkan item berdasarkan ID dari database
func GetItem(c *gin.Context) {
    id := c.Param("id")
    item, err := models.GetItem(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
        return
    }
    c.JSON(http.StatusOK, item)
}

// CreateItem menambahkan item baru ke database
func CreateItem(c *gin.Context) {
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := models.CreateItem(item); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, item)
}

// DeleteItem menghapus item berdasarkan ID dari database
func DeleteItem(c *gin.Context) {
    id := c.Param("id")
    if err := models.DeleteItem(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}

// UpdateItem memperbarui item berdasarkan ID di database
func UpdateItem(c *gin.Context) {
    id := c.Param("id")
    
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    item.ID = id
    if err := models.UpdateItem(item); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, item)
}

