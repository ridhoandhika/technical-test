package controllers

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckServer(){
	for i := 0; i < 100; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func GetHealth(c *gin.Context) {
	go CheckServer()
	fmt.Println("test ridho")

	time.Sleep(1 * time.Second)

	c.JSON(http.StatusOK, gin.H{"ststus": "OK"})
}