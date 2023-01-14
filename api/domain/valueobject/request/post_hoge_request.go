package request

type PostHogeRequest struct {
	HogeName string `uri:"hogeName" binding:"required"`
}
