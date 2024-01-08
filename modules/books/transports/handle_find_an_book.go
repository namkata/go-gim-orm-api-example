package bookview

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"local-app/common"
	appctx "local-app/component/appcontext"

	bookbiz "local-app/modules/books/biz"
	bookmodel "local-app/modules/books/models"
	bookstorage "local-app/modules/books/storages"
)

// HandleFindAnBook handles the retrieval of a book by its ID.
func HandleFindAnBook(appctx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				gin.H{"id": id},
				gin.H{"data": false},
			))
			return
		}
		db := appctx.GetDBConnection()
		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewFindBookBiz(storage)

		data, err := biz.FindAnBook(ctx.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			if errors.Is(err, bookmodel.ErrBookNotFound) {
				ctx.JSON(http.StatusNotFound, common.BadResponseFailure(
					"Book not found",
					gin.H{"id": id},
					gin.H{"data": false},
				))
				return
			}
			ctx.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				gin.H{"id": id},
				gin.H{"data": false},
			))
			return
		}

		ctx.JSON(http.StatusOK, common.ResponseOKSuccess(
			"Book found successfully",
			gin.H{"id": id},
			gin.H{"data": data},
		))
	}
}
