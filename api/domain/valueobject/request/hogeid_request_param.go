package request

type HogeIDRequestParam struct {
	HogeID uint64 `uri:"hogeID" binding:"required"`
}
