package net

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/isurusiri/services/shared/models"
	"github.com/isurusiri/services/shared/utils"
)

// GetRemote invokes a given remote endpoint and
// retuen the response.
func GetRemote(uri string) (*models.APIResponse, error) {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	remote, err := utils.PareseAPIResponse(body)

	return remote, err
}
