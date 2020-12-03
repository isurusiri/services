package net

import (
	"log"
	"net/http"
)

// GetRemote invokes a given remote endpoint and
// retuen the response.
func GetRemote(uri string) *http.Response {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}
