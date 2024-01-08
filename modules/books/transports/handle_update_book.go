package bookview

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"local-app/common"
	appctx "local-app/component/appcontext"

	bookbiz "local-app/modules/books/biz"
	bookmodel "local-app/modules/books/models"
	bookstorage "local-app/modules/books/storages"
)

// HandleUpdateAnBook handles the HTTP request to update a book's information.
func HandleUpdateAnBook(appctx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// Extract the book ID from the URL parameter
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadGateway, common.BadResponseFailure(
				err.Error(),
				gin.H{"id": id},
				gin.H{"data": false},
			))
			return
		}

		// Retrieve book data from the request body
		var bookData bookmodel.Book
		if err := ctx.ShouldBind(&bookData); err != nil {
			ctx.JSON(http.StatusBadGateway, common.BadResponseFailure(
				err.Error(),
				gin.H{"id": id},
				gin.H{"data": false},
			))
			return
		}

		// Create storage instance and business logic
		db := appctx.GetDBConnection()
		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewUpdateBookBiz(storage)

		// Update the book data using business logic
		if err := biz.UpdateBook(ctx.Request.Context(), map[string]interface{}{"id": id}, &bookData); err != nil {
			ctx.JSON(http.StatusBadGateway, common.BadResponseFailure(
				err.Error(),
				gin.H{"id": id},
				gin.H{"data": false},
			))
			return
		}

		// Return success response if book update is successful
		ctx.JSON(http.StatusOK, common.BadResponseFailure(
			"Update a book successfully",
			gin.H{"id": id},
			gin.H{"data": true},
		))
	}
}
