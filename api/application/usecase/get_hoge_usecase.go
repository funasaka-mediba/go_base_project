package usecase

import (
	"go_base_project/constant"
	"go_base_project/domain/repository"
	"go_base_project/domain/valueobject/response"
	"go_base_project/packages/customError"
	"time"

	"github.com/gin-gonic/gin"
)

// GetHogeUsecase
type GetHogeUsecase interface {
	GetHoge(ctx *gin.Context) (*response.GetHogeResponse, *customError.CustomError)
}

type getHogeUsecase struct {
	hr repository.HogeRepository
}

// NewGetHogeUsecase コンストラクタ
func NewGetHogeUsecase(hr repository.HogeRepository) GetHogeUsecase {
	return getHogeUsecase{hr}
}

func (u getHogeUsecase) GetHoge(ctx *gin.Context) (*response.GetHogeResponse, *customError.CustomError) {
	return &response.GetHogeResponse{
		Timestamp: time.Now().Format(constant.DateTimeLayout),
		Results:   &response.GetHogeResult{},
	}, nil
}
