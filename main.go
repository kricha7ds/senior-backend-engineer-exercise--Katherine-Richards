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

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "success",
	})
}

func employees(c *gin.Context) {
	var employees []Employee
	rows := DB.Find(&employees)

	if rows.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found.",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"employees": employees,
	})
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

	// Preload associated record to be displayed in the response
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
	router.GET("/", home)
	router.GET("/employees", employees)
	router.POST("/jobs", createJobs)

	router.Run(":8080")
}
