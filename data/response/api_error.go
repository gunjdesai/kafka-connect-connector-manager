package response

type ApiErrorResponse struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}
