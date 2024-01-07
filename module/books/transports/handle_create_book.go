package bookview

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	bookbiz "local-app/module/books/biz"
	bookmodel "local-app/module/books/models"
	bookstorage "local-app/module/books/storages"
)

// HandleCreateBook handles the creation of a new book entry endpoint.
func HandleCreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookItem bookmodel.Book

		// Retrieve book information from the request
		if err := ctx.ShouldBind(&bookItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Pre-process title - trim leading and trailing spaces
		bookItem.Name = strings.TrimSpace(bookItem.Name)

		// Initialize storage and business logic components
		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewCreateBookBiz(storage)

		// Attempt to create a new book entry
		if err := biz.CreateBook(ctx.Request.Context(), &bookItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Respond with the ID of the created book
		ctx.JSON(http.StatusOK, gin.H{"data": bookItem.ID})
	}
}
