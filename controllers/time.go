package controllers

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
    DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func GetTime(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"time": time.Now().Format(DDMMYYYYhhmmss), "message": "Hello World!"})
}