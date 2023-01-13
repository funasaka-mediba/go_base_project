package response

type GetHogeResponse struct {
	Timestamp string         `json:"timestamp"`
	Result    *GetHogeResult `json:"result"`
}

type GetHogeResult struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
