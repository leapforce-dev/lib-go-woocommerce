package woocommerce

// ErrorResponse stores general API error response
//
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status int               `json:"status"`
		Params map[string]string `json:"params"`
	} `json:"data"`
}
