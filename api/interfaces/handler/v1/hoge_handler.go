package v1

import (
	"github.com/gin-gonic/gin"
)

type HogeHandler interface {
	GetHoge(ctx *gin.Context)
}

type hogeHandler struct {
	// hu usecase.HogeUsecase
}

// NewHogeHandler コンストラクタ
func NewHogeHandler() HogeHandler {
	return hogeHandler{}
}

func (h hogeHandler) GetHoge(ctx *gin.Context) {
	// res, err := h.hu.GetHoge(ctx)
	// if err != nil {
	// 	log.Logger.Error("failed to [GET]/hoge", zap.String("path", ctx.Request.URL.Path), zap.Error(err.Err), zap.String("error_source", err.ErrSource))
	// 	ctx.JSON(err.Status, CreateErrorResponse(util.GetRequestIDFromHeader(ctx), err.Err, err.Code, err.RefURL))
	// 	return
	// }
	// ctx.JSON(200, res)
}
