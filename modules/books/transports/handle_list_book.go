package bookview

import (
	"net/http"

	"local-app/common"
	appctx "local-app/component/appcontext"

	"github.com/gin-gonic/gin"

	bookbiz "local-app/modules/books/biz"
	bookmodel "local-app/modules/books/models"
	bookstorage "local-app/modules/books/storages"
)

// ListBookRequest represents the combined request structure for listing books.
type ListBookRequest struct {
	common.Paging    // Embedding common.Paging for pagination
	bookmodel.Filter // Embedding bookmodel.Filter for filtering
}

// HandleListBook handles the listing of books endpoint.
func HandleListBook(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request ListBookRequest

		// Retrieve paging information from the request
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				request,
				gin.H{"data": []common.Dictionary{}, "paging": &request.Paging}),
			)
			return
		}

		// Process the paging information
		request.ProcessPaging()

		// Initialize storage and business logic components
		db := appCtx.GetDBConnection()
		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewListBookBiz(storage)

		// Retrieve a list of books based on the provided paging parameters and filter
		result, err := biz.ListBook(ctx.Request.Context(), &request.Filter, &request.Paging)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, common.BadResponseFailure(
				err.Error(),
				request,
				gin.H{"data": result, "paging": &request.Paging}),
			)
			return
		}

		// Return the fetched data and paging information in the response
		ctx.JSON(http.StatusOK, common.ResponseOKSuccess(
			"List book successfully",
			request,
			gin.H{"data": result, "paging": &request.Paging}),
		)
	}
}
