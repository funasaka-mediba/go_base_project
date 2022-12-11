package response

type GetHogeResponse struct {
	Timestamp string         `json:"timestamp"`
	Results   *GetHogeResult `json:"results"`
}

type GetHogeResult struct{}
