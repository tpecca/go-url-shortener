package main

import (
	"context"
	"go-url-shortener/db"
	"go-url-shortener/handlers"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Missing environment configuration")
	}

	dbURI := os.Getenv("MONGODB_URI")
	if dbURI == "" {
		log.Fatal("Database connection string not provided")
	}

	if err := db.SetupMongoDB(dbURI); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.MongoClient.Disconnect(context.Background())

	router := gin.Default()
	router.LoadHTMLFiles("static/form.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "form.html", nil)
	})

	router.POST("/shorten", handlers.CreateShortLink)
	router.GET("/:hash", handlers.RedirectToOriginal)

	router.Run(":4000")
}
