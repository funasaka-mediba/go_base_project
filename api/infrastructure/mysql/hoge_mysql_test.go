package mysql

import (
	"go_base_project/domain/entity/mysqlEntity"
	"go_base_project/infrastructure/adaptor"
	"go_base_project/packages/customError"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_hogeMySQL_GetHoge(t *testing.T) {
	// gin.SetMode(gin.ReleaseMode) ウェルネスのテストクラスだとReleaseModeで実行している。
	gin.SetMode(gin.TestMode)                               // これがないとginはdebugモードで起動する。なんとなくReleaseModeよりTestModeのほうが正しい気がする。
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder()) // テスト用に適当なcontextを作成
	db, _ := adaptor.ReadDb(ctx)
	tests := map[string]struct {
		input  func() (hogeID int)
		output func(*testing.T, *mysqlEntity.Hoge, *customError.CustomError)
	}{
		// TODO: Add test cases.
		"success": {
			input: func() (hogeID int) {
				return 0
			},
			output: func(t *testing.T, h *mysqlEntity.Hoge, err *customError.CustomError) {
				assert.Nilf(t, err, "[success] err: %+v", err)
				assert.Equal(t, h.ID, uint64(0), "[success] h.ID: %+v", h.ID)
				assert.Equal(t, h.Name, "hogehogetesttest", "[success] h.Name: %+v", h.Name)
			},
		},
	}
	h := &hogeMySQL{
		db: db,
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			hogeID := tt.input()
			exist, err := h.GetHoge(ctx, hogeID)
			tt.output(t, exist, err)
		})
	}
}
