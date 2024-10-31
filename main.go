package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Gender string `json:"gender" gorm:"not null"`
}

type Job struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Department string    `json:"department" gorm:"not null"`
	JobTitle   string    `json:"job_title" gorm:"not null"`
	EmployeeID uint      `json:"employee_id" gorm:"unique"`
	Employee   Employee  `json:"employee"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

var DB *gorm.DB

func dbConnect() {
	db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database.")
	}

	db.AutoMigrate(&Employee{}, &Job{})

	DB = db
}

func createJobs(c *gin.Context) {
	var jobs []Job

	if err := c.ShouldBindJSON(&jobs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	rows := DB.CreateInBatches(jobs, 100)
	if rows.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error.",
		})
		return
	}

	if err := DB.Preload("Employee").Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching jobs.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"jobs": jobs,
	})
}

func main() {
	dbConnect()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/jobs", createJobs)
	}

	router.Run(":8080")
}
