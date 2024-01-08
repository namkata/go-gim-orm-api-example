package bookroute

import (
	appctx "local-app/component/appcontext"
	bookview "local-app/modules/books/transports"

	"github.com/gin-gonic/gin"
)

func Routes(v1 *gin.RouterGroup, appCtx appctx.AppContext) {
	books := v1.Group("/books")
	{
		books.GET("", bookview.HandleListBook(appCtx))
		books.POST("", bookview.HandleCreateBook(appCtx))
		books.GET("/:id", bookview.HandleFindAnBook(appCtx))
		books.PUT("/:id", bookview.HandleUpdateAnBook(appCtx))
		books.DELETE("/:id", bookview.HandleDeleteAnBook(appCtx))
	}
}
