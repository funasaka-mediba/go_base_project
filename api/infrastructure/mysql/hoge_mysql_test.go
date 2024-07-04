package mysql

import (
	"context"
	"go_base_project/domain/entity/mysqlEntity"
	"go_base_project/infrastructure/adaptor"
	"go_base_project/packages/customError"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHogeMySQL_GetHoges(t *testing.T) {
	db, _ := adaptor.WriteDb(context.Background())

	tests := map[string]struct {
		input  func()
		output func(*testing.T, []mysqlEntity.Hoge, *customError.CustomError)
	}{
		"success": {
			input: func() {
			},
			output: func(t *testing.T, h []mysqlEntity.Hoge, err *customError.CustomError) {
				assert.Nil(t, err)
				assert.Equal(t, 8, len(h))
				assert.Equal(t, uint64(1), h[0].ID)
				assert.Equal(t, "kaguya", h[0].Name)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.input()
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			h, err := NewHogeMySQL(db).GetHoges(ctx)
			tt.output(t, h, err)
		})
	}
}
