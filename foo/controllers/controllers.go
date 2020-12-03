package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/isurusiri/services/shared/models"
)

var apiResponse models.APIResponse

// FooController handles the root url of the bar service
// Request path: /
func FooController(w http.ResponseWriter, r *http.Request) {
	apiResponse.APIVersion = "v1"
	apiResponse.APIName = "foo"
	apiResponse.AppVersion = "v0.0.1"
	apiResponse.ResourcePath = "/"

	// get remote

	json.NewEncoder(w).Encode(apiResponse)
}
