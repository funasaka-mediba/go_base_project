package constant

import "go_base_project/packages/env"

type ErrorCode string

const (
	// 400系エラーコード
	GBP4040 ErrorCode = "GBP4040" // 指定されたコンテンツが存在しない

	// 500系エラーコード
	GBP5000 ErrorCode = "GBP5000" // サーバー内部エラー
)

// GetRefURL .
func GetRefURL(path string) string {
	return env.Env().HostURL + path
}
