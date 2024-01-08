package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	appctx "local-app/component/appcontext"

	bookmodel "local-app/modules/books/models"
	bookroute "local-app/routes/books"
)

var db *gorm.DB // Database connection

func main() {
	// Load the environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to SQLite database
	var err error

	// // Read environment variables
	// dbHost := os.Getenv("DB_HOST")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		// Example usage for other databases
		// DSN: fmt.Sprintf("%s:%s@tcp(%s)/dbname", dbUser, dbPassword, dbHost),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&bookmodel.Book{}); err != nil {
		log.Fatal(err)
	}

	appCtx := appctx.NewAppContext(db)
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	routeGroups := router.Group("/api/v1")
	{
		// Define a health check endpoint
		routeGroups.GET("/health-check", heathCheck)
		bookroute.Routes(routeGroups, appCtx)
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
