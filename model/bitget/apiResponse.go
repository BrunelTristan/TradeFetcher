package bitget

type ApiResponse struct {
	Code               string
	Message            string `json:"msg"`
	RequestedTimestamp int64
}
