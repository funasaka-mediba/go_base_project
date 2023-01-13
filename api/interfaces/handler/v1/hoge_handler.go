package v1

import (
	"go_base_project/application/usecase"
	"go_base_project/packages/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HogeHandler interface {
	GetHoges(ctx *gin.Context)
	GetHoge(ctx *gin.Context)
}

type hogeHandler struct {
	hu usecase.HogeUsecase
}

// NewHogeHandler コンストラクタ
func NewHogeHandler(hu usecase.HogeUsecase) HogeHandler {
	return hogeHandler{hu}
}

func (h hogeHandler) GetHoges(ctx *gin.Context) {
	res, err := h.hu.GetHoges(ctx)
	if err != nil {
		log.Logger.Error("failed to [GET]/hoge", zap.String("path", ctx.Request.URL.Path), zap.Error(err.Err), zap.String("error_source", err.ErrSource))
		ctx.JSON(err.Status, CreateErrorResponse(err.Err, err.Code, err.RefURL))
		return
	}
	ctx.JSON(200, res)
}

func (h hogeHandler) GetHoge(ctx *gin.Context) {
	res, err := h.hu.GetHoge(ctx)
	if err != nil {
		log.Logger.Error("failed to [GET]/hoge", zap.String("path", ctx.Request.URL.Path), zap.Error(err.Err), zap.String("error_source", err.ErrSource))
		ctx.JSON(err.Status, CreateErrorResponse(err.Err, err.Code, err.RefURL))
		return
	}
	ctx.JSON(200, res)
}
