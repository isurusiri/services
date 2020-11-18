package controllers

import (
	"fmt"
	"net/http"
)

// BarController handles the root url of the bar service
// Request path: /
func BarController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Bar!")

}
