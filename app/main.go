package main

import (
	"database/sql"
	"grade-calculator/views"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/data.db")

	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {
	// Initialize SQLite database
	db, err := initDB()
	if err != nil {
		log.Fatal("Error initializing the database:", err)
	}
	defer db.Close()

	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		render(c, http.StatusOK, views.MakeHome("Mateo's Cool", "Felipe"))
	})

	r.Run(":9090")
}
