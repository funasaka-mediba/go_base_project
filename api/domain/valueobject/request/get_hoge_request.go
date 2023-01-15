package request

type GetHogeRequest struct {
	HogeID uint64 `uri:"hoge_id" binding:"required"`
}
