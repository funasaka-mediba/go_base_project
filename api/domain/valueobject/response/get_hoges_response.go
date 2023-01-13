package response

type GetHogesResponse struct {
	Timestamp string           `json:"timestamp"`
	Results   *GetHogesResults `json:"results"`
}

type GetHogesResults struct {
	List []GetHogesResult `json:"list"`
}

type GetHogesResult struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
