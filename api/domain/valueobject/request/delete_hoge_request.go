package request

type DeleteHogeRequest struct {
	HogeID uint64 `uri:"hoge_id" binding:"required"`
}
