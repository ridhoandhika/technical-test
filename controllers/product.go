package controllers

import (
    "rest-test/models"
    "net/http"
    

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

type CreateProductInput struct {
    Name 				string `json:"Name"`
    Description  		string `json:"Description"`
}

type UpdateProductInput struct {
    Name 				string `json:"Name"`
    Description  		string `json:"Description"`
}

// Method => GET /products
// Return all product
func GetProduct(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var product []models.Product
    db.Find(&product)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// Method => POST /products
// Insert product baru
func CreateTask(c *gin.Context) {
    // validasi inputan/body
    var input CreateProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Insert products
    product := models.Product{Product_Name: input.Name, Product_Description: input.Description}

    db := c.MustGet("db").(*gorm.DB)
    db.Create(&product)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// Method =>  GET /products/:id
// Cari product berdasarkan ID
func FindProduct(c *gin.Context) { 
   
    var product models.Product

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("product_id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// Method => PATCH /products/:id
// Update a products
func UpdateProduct(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    
    var product models.Product
    if err := db.Where("product_id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
        return
    }

    // Validasi input/body
    var input UpdateProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Product
    updatedInput.Product_Name = input.Name
    updatedInput.Product_Description = input.Description

    db.Model(&product).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// Method => DELETE /products/:id
// Delete products berdasarkan ID
func DeleteProduct(c *gin.Context) {
    
    db := c.MustGet("db").(*gorm.DB)
    var product models.Product
    if err := db.Where("product_id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
        return
    }

    db.Delete(&product)

    c.JSON(http.StatusOK, gin.H{"data": true})
}