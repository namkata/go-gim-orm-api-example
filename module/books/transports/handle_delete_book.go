package bookview

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	bookbiz "local-app/module/books/biz"
	bookstorage "local-app/module/books/storages"
)

// HandleDeleteAnBook handles the deletion of a book by ID.
func HandleDeleteAnBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the book ID from the request URL parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a new storage instance
		storage := bookstorage.NewDBStorage(db)

		// Create a business logic instance for deleting books
		biz := bookbiz.NewDeleteBookBiz(storage)

		// Attempt to delete the book using the provided ID
		if err := biz.DeleteBook(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return success message if deletion is successful
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
