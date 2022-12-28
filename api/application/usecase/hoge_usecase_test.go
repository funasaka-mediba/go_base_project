package usecase

import (
	"go_base_project/domain/repository"
	"go_base_project/domain/valueobject/response"
	"go_base_project/packages/customError"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_hogeUsecase_GetHoge(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := map[string]struct {
		output func(*testing.T, *response.GetHogeResponse, *customError.CustomError)
	}{
		// TODO: Add test cases.
		"success": {
			output: func(t *testing.T, ghr *response.GetHogeResponse, err *customError.CustomError) {
				assert.Nilf(t, err, "[success] err: %+v", err)
				assert.Equal(t, ghr.Results.ID, uint64(0), "[success]ghr.Results.ID: %+v", ghr.Results.ID)
				assert.Equal(t, ghr.Results.Name, "hogehogetesttest", "[success]ghr.Results.Name: %+v", ghr.Results.Name)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			u := hogeUsecase{
				hr: repository.NewHogeRepository(),
			}
			res, err := u.GetHoge(ctx)
			tt.output(t, res, err)
		})
	}
}
