package models
import (
"fmt"
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)
// Configurasi koneksi database mysql
func SetupDB() *gorm.DB {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "learnGo"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	
	if err != nil {
	panic(err.Error())
	}
	return db
}