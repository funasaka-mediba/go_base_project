package repository

import (
	"go_base_project/domain/entity/mysqlEntity"
	"go_base_project/infrastructure/adaptor"
	"go_base_project/infrastructure/mysql"
	"go_base_project/packages/customError"

	"github.com/gin-gonic/gin"
)

type HogeRepository interface {
	GetHoge(ctx *gin.Context, hogeID int) (*mysqlEntity.Hoge, *customError.CustomError)
}

type hogeRepository struct {
}

func NewHogeRepository() HogeRepository {
	return hogeRepository{}
}

func (r hogeRepository) GetHoge(ctx *gin.Context, hogeID int) (*mysqlEntity.Hoge, *customError.CustomError) {
	db, err := adaptor.WriteDb(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewHogeMySQL(db).GetHoge(ctx, hogeID)
}
