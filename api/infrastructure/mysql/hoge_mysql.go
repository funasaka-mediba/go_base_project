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
	InsertHoge(ctx *gin.Context, name string) (int64, *customError.CustomError)
	DeleteHoge(ctx *gin.Context, hogeID uint64) *customError.CustomError
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

func (m *hogeMySQL) InsertHoge(ctx *gin.Context, name string) (int64, *customError.CustomError) {
	query := strings.Join([]string{
		"INSERT INTO `hoge` (`name`) VALUES (?)",
	}, " ")

	result, err := m.db.Exec(ctx, query, name)
	if err != nil {
		return 0, customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
	}

	return insertID, nil
}

func (m *hogeMySQL) DeleteHoge(ctx *gin.Context, hogeID uint64) *customError.CustomError {
	query := strings.Join([]string{
		"DELETE FROM `hoge`",
		"WHERE",
		"	`id` = ?",
	}, " ")

	if _, err := m.db.Exec(ctx, query, hogeID); err != nil {
		return customError.NewErr(err, constant.GBP5000, customError.Error, 400, "")
	}
	return nil
}
