package main

import "fmt"

var app App

func main() {
	fmt.Println("Starting bar api.")
	app.Initialize()
	// env arg?
	app.Run(":8080")
}
