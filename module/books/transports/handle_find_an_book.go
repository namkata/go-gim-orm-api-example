package bookview

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	bookbiz "local-app/module/books/biz"
	bookmodel "local-app/module/books/models"
	bookstorage "local-app/module/books/storages"
)

// HandleFindAnBook handles the retrieval of a book by its ID.
func HandleFindAnBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := bookstorage.NewDBStorage(db)
		biz := bookbiz.NewFindBookBiz(storage)

		data, err := biz.FindAnBook(ctx.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			if errors.Is(err, bookmodel.ErrBookNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
				return
			}
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
