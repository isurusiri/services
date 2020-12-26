package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/isurusiri/services/foo/context"
	"github.com/isurusiri/services/foo/net"
	"github.com/isurusiri/services/shared/models"
)

var apiResponse models.APIResponse
var appContext context.AppContext

// FooController handles the root url of the bar service
// Request path: /
func FooController(w http.ResponseWriter, r *http.Request) {
	apiResponse.APIVersion = "v1"
	apiResponse.APIName = "foo"
	apiResponse.AppVersion = "v0.0.1"
	apiResponse.ResourcePath = "/"

	// get remote
	remote, err := net.GetRemote(context.GetRemoteHost())
	if err != nil {
		log.Println(err)
	}
	apiResponse.Remote = remote

	json.NewEncoder(w).Encode(apiResponse)
}
