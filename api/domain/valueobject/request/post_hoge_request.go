package request

type PostHogeRequest struct {
	HogeName string `uri:"hoge_name" binding:"required"`
}
