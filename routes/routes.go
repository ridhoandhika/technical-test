package routes

import (
    "rest-test/controllers"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })
    r.GET("/products", controllers.GetProduct)
    r.POST("/products", controllers.CreateTask)
    r.GET("/products/:id", controllers.FindProduct)
    r.PATCH("/products/:id", controllers.UpdateProduct)
    r.DELETE("products/:id", controllers.DeleteProduct)

	r.GET("/health", controllers.GetHealth)
	r.GET("/time", controllers.GetTime)
    return r
}