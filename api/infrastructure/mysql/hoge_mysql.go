package mysql

import (
	"go_base_project/constant"
	"go_base_project/domain/entity/mysqlEntity"
	"go_base_project/infrastructure/adaptor"
	"go_base_project/packages/customError"
	"strings"

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
	query := strings.Join([]string{
		"SELECT",
		"	`id`,",
		"	`name`",
		"FROM",
		"	`hoge`",
		"WHERE",
		"	`id` = ?",
	}, " ")

	var h mysqlEntity.Hoge
	if err := m.db.Get(ctx, &h, query, hogeID); err != nil {
		return nil, customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
	}
	return &h, nil
}
