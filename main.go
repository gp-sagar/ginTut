package main

import (
	"ginTut/middleware"
	"ginTut/models"
	"ginTut/router"
	"log"
	"os"

	_ "ginTut/docs" // Import the docs package to initialize the swagger docs

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Tutorial API
// @version 1.0
// @description ginTut Api Docs

// @contact.name API Support
// @contact.url localhost:8080
// @contact.email sagarsingh@polarisgrids.com

// @host localhost:8080
// @BasePath /
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	dsn := "host=" + os.Getenv("DATABASE_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	err = models.InitDB(dsn)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.NoRoute(middleware.NotFoundHandler())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize routes
	router.InitializeRoutes(r)

	r.Run(":8080")
}
