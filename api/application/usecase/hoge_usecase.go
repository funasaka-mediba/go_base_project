package usecase

import (
	"go_base_project/constant"
	"go_base_project/domain/repository"
	"go_base_project/domain/valueobject/response"
	"go_base_project/packages/customError"
	"time"

	"github.com/gin-gonic/gin"
)

// HogeUsecase
type HogeUsecase interface {
	GetHoge(ctx *gin.Context) (*response.GetHogeResponse, *customError.CustomError)
}

type hogeUsecase struct {
	hr repository.HogeRepository
}

// NewGetHogeUsecase コンストラクタ
func NewHogeUsecase(hr repository.HogeRepository) HogeUsecase {
	return hogeUsecase{hr}
}

func (u hogeUsecase) GetHoge(ctx *gin.Context) (*response.GetHogeResponse, *customError.CustomError) {
	hoge, err := u.hr.GetHoge(ctx, 0)
	if err != nil {
		return nil, err
	}

	return &response.GetHogeResponse{
		Timestamp: time.Now().Format(constant.DateTimeLayout),
		Results: &response.GetHogeResult{
			ID:   hoge.ID,
			Name: hoge.Name,
		},
	}, nil
}
