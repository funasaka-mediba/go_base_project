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
	// TODO:query作成
	// TODO:DBからselectの処理を後で書く

	var h mysqlEntity.Hoge

	// TODO:あとでDBからちゃんと取得するようにする
	h.ID = uint64(hogeID)
	h.Name = "hogehogetesttest"
	return &h, nil
}
