package repository

import (
	"go_base_project/domain/entity/mysqlEntity"
	"go_base_project/infrastructure/adaptor"
	"go_base_project/infrastructure/mysql"
	"go_base_project/packages/customError"

	"github.com/gin-gonic/gin"
)

type HogeRepository interface {
	GetHoges(ctx *gin.Context) ([]mysqlEntity.Hoge, *customError.CustomError)
	GetHoge(ctx *gin.Context, hogeID uint64) (*mysqlEntity.Hoge, *customError.CustomError)
	InsertHoge(ctx *gin.Context, name string) (int64, *customError.CustomError)
}

type hogeRepository struct {
}

func NewHogeRepository() HogeRepository {
	return hogeRepository{}
}

func (r hogeRepository) GetHoges(ctx *gin.Context) ([]mysqlEntity.Hoge, *customError.CustomError) {
	db, err := adaptor.WriteDb(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewHogeMySQL(db).GetHoges(ctx)
}

func (r hogeRepository) GetHoge(ctx *gin.Context, hogeID uint64) (*mysqlEntity.Hoge, *customError.CustomError) {
	db, err := adaptor.WriteDb(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewHogeMySQL(db).GetHoge(ctx, hogeID)
}

func (r hogeRepository) InsertHoge(ctx *gin.Context, name string) (int64, *customError.CustomError) {
	db, err := adaptor.WriteDb(ctx)
	if err != nil {
		return 0, err
	}
	return mysql.NewHogeMySQL(db).InsertHoge(ctx, name)
}
