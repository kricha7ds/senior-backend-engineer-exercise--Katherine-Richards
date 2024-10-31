package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func dbConnect() {
	db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}

	DB = db
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "success",
	})
}

func main() {
	dbConnect()

	router := gin.Default()
	router.GET("/", home)

	router.Run(":8080")
}
