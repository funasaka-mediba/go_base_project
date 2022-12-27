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
	gin.SetMode(gin.ReleaseMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
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
