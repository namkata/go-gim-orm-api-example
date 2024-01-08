package bookview

import (
	"local-app/common"
	appctx "local-app/component/appcontext"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	bookbiz "local-app/modules/books/biz"
	bookstorage "local-app/modules/books/storages"
)

// HandleDeleteAnBook handles the deletion of a book by ID.
func HandleDeleteAnBook(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get the book ID from the request URL parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				gin.H{"id": id},
				gin.H{"data": false},
			))
			return
		}

		// Create a new storage instance
		db := appctx.GetDBConnection()
		storage := bookstorage.NewDBStorage(db)

		// Create a business logic instance for deleting books
		biz := bookbiz.NewDeleteBookBiz(storage)

		// Attempt to delete the book using the provided ID
		if err := biz.DeleteBook(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				gin.H{"id": id},
				gin.H{"data": false},
			))
			return
		}

		// Return success message if deletion is successful
		c.JSON(http.StatusOK, common.ResponseOKSuccess(
			"Delete a book successfully",
			gin.H{"id": id},
			gin.H{"data": true},
		))
	}
}
