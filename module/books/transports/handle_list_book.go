package bookview

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	bookbiz "local-app/module/books/biz"
	bookmodel "local-app/module/books/models"
	bookstorage "local-app/module/books/storages"
)

// HandleListBook handles the listing of books endpoint.
func HandleListBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging bookmodel.DataPaging

		// Retrieve paging information from the request
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Process the paging information
		paging.ProcessPaging()

		// Initialize storage and business logic components
		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewListBookBiz(storage)

		// Retrieve a list of books based on the provided paging parameters
		result, err := biz.ListBook(ctx.Request.Context(), nil, &paging)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		// Return the fetched data and paging information in the response
		ctx.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
