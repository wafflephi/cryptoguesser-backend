package responses

type AuthResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
