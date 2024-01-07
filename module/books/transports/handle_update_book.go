package bookview

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	bookbiz "local-app/module/books/biz"
	bookmodel "local-app/module/books/models"
	bookstorage "local-app/module/books/storages"
)

// HandleUpdateAnBook handles the HTTP request to update a book's information.
func HandleUpdateAnBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract the book ID from the URL parameter
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		// Retrieve book data from the request body
		var bookData bookmodel.Book
		if err := ctx.ShouldBind(&bookData); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		// Create storage instance and business logic
		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewUpdateBookBiz(storage)

		// Update the book data using business logic
		if err := biz.UpdateBook(ctx.Request.Context(), map[string]interface{}{"id": id}, &bookData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return success response if book update is successful
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
