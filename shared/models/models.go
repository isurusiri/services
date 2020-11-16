package models

// APIResponse represents a generic API response object.
type APIResponse struct {
	APIVersion   string       `json:"apiVersion"`
	AppVersion   string       `json:"appVersion"`
	ResourcePath string       `json:"resourcePath"`
	APIName      string       `json:"apiName"`
	Remote       *APIResponse `json:"remote"`
}
