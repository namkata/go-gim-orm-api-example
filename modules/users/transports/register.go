package userview

import (
	"local-app/common"
	appctx "local-app/component/appcontext"
	userbiz "local-app/modules/users/biz"
	usermodel "local-app/modules/users/models"
	userstorage "local-app/modules/users/storages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetDBConnection()
		var data usermodel.User

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
		}

		storage := userstorage.NewDBStorage(db)
		biz := userbiz.NewCreateUserBiz(storage)

		if err := biz.CreateUser(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.BadResponseFailure(
				err.Error(),
				data,
				gin.H{"data": false},
			))
			return
		}

		ctx.JSON(http.StatusOK, common.ResponseOKSuccess(
			"Create a user successfully",
			data,
			gin.H{"data": data.Id},
		))
	}
}
