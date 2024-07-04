package mysql

import (
	"os"
	"testing"

	"go_base_project/packages/test"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// 初期化処理
	test.InitTest("../../config/test.env")
	gin.SetMode(gin.ReleaseMode)

	code := m.Run()
	// 後処理
	os.Exit(code)
}
