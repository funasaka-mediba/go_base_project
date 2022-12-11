package middleware

import (
	"go_base_project/constant"
	"go_base_project/domain/valueobject/request"
	v1 "go_base_project/interfaces/handler/v1"

	"github.com/gin-gonic/gin"
)

// BindBaseRequestHeader .
func BindBaseRequestHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := request.BaseRequestHeader{}
		if err := ctx.ShouldBindHeader(&header); err != nil {
			ctx.JSON(400,
				v1.CreateErrorResponse(
					err,
					constant.GBP4100,
					"",
				))
			ctx.Abort()
		}
	}
}
