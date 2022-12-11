package mysql

import (
	"go_base_project/domain/entity/mysqlEntity"
	"go_base_project/infrastructure/adaptor"
	"go_base_project/packages/customError"

	"github.com/gin-gonic/gin"
)

type HogeMySQL interface {
	GetHoge(ctx *gin.Context, hogeID int) (*mysqlEntity.Hoge, *customError.CustomError)
}

type hogeMySQL struct {
	db adaptor.DBAdaptor
}

// NewHogeMySQL コンストラクタ
func NewHogeMySQL(db adaptor.DBAdaptor) HogeMySQL {
	return &hogeMySQL{db: db}
}

func (m *hogeMySQL) GetHoge(ctx *gin.Context, hogeID int) (*mysqlEntity.Hoge, *customError.CustomError) {
	// query作成

	var h mysqlEntity.Hoge

	// DBからselect

	return &h, nil
}
