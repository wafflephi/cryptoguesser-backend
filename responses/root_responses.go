package responses

type VersionResponse struct {
	Version      string                 `json:"version"`
	Version_name string                 `json:"version_name"`
	Error_Data   map[string]interface{} `json:"error_data"`
}

type TransactionResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
