package models

import (
    "time"
)

type Product struct {
    Product_Id 				        int     `json:"Produt_Id" gorm:"primary_key"`
    Product_Name 			        string  `json:"Produt_Name"`
    Product_Description              string   `json:"Produt_Description"`
    CreatedAt  				        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  				        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}