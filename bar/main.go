package main

import (
	"fmt"
	"os"

	"github.com/isurusiri/services/bar/context"
)

var app App
var appContext context.AppContext

func initContext() {
	appContext := context.New(os.Getenv("BAR_RUNNING_PORT"))
	fmt.Printf("AppContext is:%v\n", appContext)
}

func main() {
	fmt.Println("Starting bar api.")
	initContext()
	app.Initialize()

	app.Run(fmt.Sprintf(":%v", context.GetRunningPort()))
}
