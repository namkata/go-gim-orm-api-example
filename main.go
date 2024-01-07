package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	bookmodel "local-app/module/books/models"
	bookview "local-app/module/books/transports"
)

var db *gorm.DB // Database connection

func main() {
	// Connect to SQLite database
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&bookmodel.Book{}); err != nil {
		log.Fatal(err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	routeGroups := router.Group("/api/v1")
	{
		// Define a health check endpoint
		routeGroups.GET("/health-check", heathCheck)
		bookGroup := routeGroups.Group("/books")
		{
			bookGroup.GET("", bookview.HandleListBook(db))
			bookGroup.POST("", bookview.HandleCreateBook(db))
			bookGroup.GET("/:id", bookview.HandleFindAnBook(db))
			bookGroup.PUT("/:id", bookview.HandleUpdateAnBook(db))
			bookGroup.DELETE("/:id", bookview.HandleDeleteAnBook(db))
		}
	}

	// Start server
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}

// heathCheck handles the server status check
func heathCheck(c *gin.Context) {
	// Respond with a JSON message indicating server status
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is running!",
	})
}
