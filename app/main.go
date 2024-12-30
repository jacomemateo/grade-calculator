package main

import (
	"fmt"
	"grade-calculator/views"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Grade struct {
	GradeValue int32
	GradeName  string
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Grade{})

	return db, err
}

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {
	// Initialize SQLite database
	db, err := initDB()
	if err != nil {
		log.Fatal("Error initalizing the database: ", err)
	}

	// Create http server
	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		render(c, http.StatusOK, views.MakeHome("Grade Calculator"))
	})

	r.POST("/submit", func(c *gin.Context) {
		// Parse form data
		gradeName := c.PostForm("gradeName")
		gradeValue := c.PostForm("gradeValue")

		// Convert gradeValue to integer
		var gradeValueInt int32
		_, err := fmt.Sscanf(gradeValue, "%d", &gradeValueInt)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid grade value")
			return
		}

		// Create a new Grade and save it to the database
		grade := Grade{
			GradeName:  gradeName,
			GradeValue: gradeValueInt,
		}

		if result := db.Create(&grade); result.Error != nil {
			c.String(http.StatusInternalServerError, "Failed to save grade")
			return
		}
	})

	r.Run(":9090")
}
