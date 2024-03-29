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
	GetHoges(ctx *gin.Context) (*response.GetHogesResponse, *customError.CustomError)
	GetHoge(ctx *gin.Context) (*response.GetHogeResponse, *customError.CustomError)
	PostHoge(ctx *gin.Context) (*response.PostHogeResponse, *customError.CustomError)
	DeleteHoge(ctx *gin.Context) (*response.DeleteHogeResponse, *customError.CustomError)
}

type hogeUsecase struct {
	hr repository.HogeRepository
}

// NewGetHogeUsecase コンストラクタ
func NewHogeUsecase(hr repository.HogeRepository) HogeUsecase {
	return hogeUsecase{hr}
}

func (u hogeUsecase) GetHoges(ctx *gin.Context) (*response.GetHogesResponse, *customError.CustomError) {
	hoges, err := u.hr.GetHoges(ctx)
	if err != nil {
		return nil, err
	}

	// もっと他にlistに詰める情報が必要ならgenerateResponseを関数として切り出す
	var list []response.GetHogesResult
	for i := range hoges {
		getHogeResult := response.GetHogesResult{
			ID:   hoges[i].ID,
			Name: hoges[i].Name,
		}
		list = append(list, getHogeResult)
	}

	return &response.GetHogesResponse{
		Timestamp: time.Now().Format(constant.DateTimeLayout),
		Results: &response.GetHogesResults{
			List: list,
		},
	}, nil
}

func (u hogeUsecase) GetHoge(ctx *gin.Context) (*response.GetHogeResponse, *customError.CustomError) {
	hogeID, _ := ctx.Get("hogeID")
	hoge, err := u.hr.GetHoge(ctx, hogeID.(uint64))
	if err != nil {
		return nil, err
	}

	return &response.GetHogeResponse{
		Timestamp: time.Now().Format(constant.DateTimeLayout),
		Result: &response.GetHogeResult{
			ID:   hoge.ID,
			Name: hoge.Name,
		},
	}, nil
}

func (u hogeUsecase) PostHoge(ctx *gin.Context) (*response.PostHogeResponse, *customError.CustomError) {
	hogeName, _ := ctx.Get("hogeName")
	id, err := u.hr.InsertHoge(ctx, hogeName.(string))
	if err != nil {
		return nil, err
	}

	return &response.PostHogeResponse{
		Timestamp: time.Now().Format(constant.DateTimeLayout),
		Result: &response.PostHogeResult{
			ID: id,
		},
	}, nil
}

func (u hogeUsecase) DeleteHoge(ctx *gin.Context) (*response.DeleteHogeResponse, *customError.CustomError) {
	hogeID, _ := ctx.Get("hogeID")
	if err := u.hr.DeleteHoge(ctx, hogeID.(uint64)); err != nil {
		return nil, err
	}

	return &response.DeleteHogeResponse{
		Timestamp: time.Now().Format(constant.DateTimeLayout),
		Result: &response.DeleteHogeResult{
			Delete: true,
		},
	}, nil
}
