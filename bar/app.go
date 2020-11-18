package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/isurusiri/services/bar/controllers"
)

// App holds the route instance of the API and
// exposes a set of function to initialize and
// interact with it.
type App struct {
	router *mux.Router
}

// Initialize , initializes the API routes.
func (a *App) Initialize() {
	a.router = mux.NewRouter().StrictSlash(true)

	// register routes.
	a.registerRoutes()
	fmt.Println("Initialized the router.")
}

// Registers routes of the API
func (a *App) registerRoutes() {
	a.router.HandleFunc("/", controllers.BarController)
}

// Run the API.
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.router))
}
