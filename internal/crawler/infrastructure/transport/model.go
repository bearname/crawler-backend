package transport

type errorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}
