package response

type GetHogeResponse struct {
	Timestamp string         `json:"timestamp"`
	Results   *GetHogeResult `json:"results"`
}

type GetHogeResult struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
