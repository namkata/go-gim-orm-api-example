package userroute

import (
	appctx "local-app/component/appcontext"
	userview "local-app/modules/users/transports"

	"github.com/gin-gonic/gin"
)

func Routes(v1 *gin.RouterGroup, appCtx appctx.AppContext) {
	v1.POST("/register", userview.Register(appCtx))

	// user := v1.Group("/users")
	// {
	// 	books.GET("", bookview.HandleListBook(appCtx))
	// 	books.POST("", bookview.HandleCreateBook(appCtx))
	// 	books.GET("/:id", bookview.HandleFindAnBook(appCtx))
	// 	books.PUT("/:id", bookview.HandleUpdateAnBook(appCtx))
	// 	books.DELETE("/:id", bookview.HandleDeleteAnBook(appCtx))
	// }
}
