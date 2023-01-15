package response

type DeleteHogeResponse struct {
	Timestamp string            `json:"timestamp"`
	Result    *DeleteHogeResult `json:"result"`
}

type DeleteHogeResult struct {
	Delete bool `json:"delete"`
}
