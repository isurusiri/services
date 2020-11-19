package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/isurusiri/services/shared/models"
)

var apiResponse models.APIResponse

// BarController handles the root url of the bar service
// Request path: /
func BarController(w http.ResponseWriter, r *http.Request) {
	apiResponse.APIVersion = "v1"
	apiResponse.APIName = "bar"
	apiResponse.AppVersion = "v0.0.1"
	apiResponse.ResourcePath = "/"

	json.NewEncoder(w).Encode(apiResponse)
}
