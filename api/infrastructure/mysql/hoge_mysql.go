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
	GetHoges(ctx *gin.Context) ([]mysqlEntity.Hoge, *customError.CustomError)
	GetHoge(ctx *gin.Context, hogeID uint64) (*mysqlEntity.Hoge, *customError.CustomError)
}

type hogeMySQL struct {
	db adaptor.DBAdaptor
}

// NewHogeMySQL コンストラクタ
func NewHogeMySQL(db adaptor.DBAdaptor) HogeMySQL {
	return &hogeMySQL{db: db}
}

func (m *hogeMySQL) GetHoges(ctx *gin.Context) ([]mysqlEntity.Hoge, *customError.CustomError) {
	query := strings.Join([]string{
		"SELECT",
		"	`id`,",
		"	`name`",
		"FROM",
		"	`hoge`",
	}, " ")

	var h []mysqlEntity.Hoge
	if err := m.db.Select(ctx, &h, query); err != nil {
		return nil, customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
	}
	return h, nil
}

func (m *hogeMySQL) GetHoge(ctx *gin.Context, hogeID uint64) (*mysqlEntity.Hoge, *customError.CustomError) {
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
