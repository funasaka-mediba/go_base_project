package response

type PostHogeResponse struct {
	Timestamp string          `json:"timestamp"`
	Result    *PostHogeResult `json:"result"`
}

type PostHogeResult struct {
	ID int64 `json:"id"`
}
