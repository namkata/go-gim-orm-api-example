package bookview

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"local-app/common"
	appctx "local-app/component/appcontext"

	bookbiz "local-app/modules/books/biz"
	bookmodel "local-app/modules/books/models"
	bookstorage "local-app/modules/books/storages"
)

// HandleCreateBook handles the creation of a new book entry endpoint.
func HandleCreateBook(appctx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var bookItem bookmodel.Book

		// Retrieve book information from the request
		if err := ctx.ShouldBind(&bookItem); err != nil {
			ctx.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				bookItem,
				gin.H{"data": false},
			))
			return
		}

		// Pre-process title - trim leading and trailing spaces
		bookItem.Name = strings.TrimSpace(bookItem.Name)

		// Initialize storage and business logic components
		db := appctx.GetDBConnection()
		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewCreateBookBiz(storage)

		// Attempt to create a new book entry
		if err := biz.CreateBook(ctx.Request.Context(), &bookItem); err != nil {
			ctx.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				bookItem,
				gin.H{"data": false},
			))
			return
		}

		// Respond with the ID of the created book
		ctx.JSON(http.StatusOK, common.ResponseOKSuccess(
			"Create a book successfully",
			bookItem,
			gin.H{"data": bookItem.ID},
		))
	}
}
