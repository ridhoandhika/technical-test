package main

import (
      
    "rest-test/models"
    "rest-test/routes"
)

func main() {
    db := models.SetupDB()

    //lakukan migration database
    db.AutoMigrate(&models.Product{})

    r := routes.SetupRoutes(db)
    r.Run()
}