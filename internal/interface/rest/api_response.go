package rest

type APIResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Errors  []APIError  `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
	Reason  string      `json:"reason,omitempty"`
}
