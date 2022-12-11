package request

// BaseRequestHeader 全APIで使用する共通のリクエストヘッダー
type BaseRequestHeader struct {
	CacheControl string `header:"Cache-Control" binding:"required"`
}
