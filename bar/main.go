package main

import (
	"fmt"
	"os"

	"github.com/isurusiri/services/bar/context"
)

var app App
var appContext context.AppContext

func initContext() {
	appContext := context.New()
	appContext.SetRunningPort(os.Getenv("BAR_RUNNING_PORT"))
}

func main() {
	fmt.Println("Starting bar api.")
	initContext()
	app.Initialize()

	app.Run(fmt.Sprintf(":%v", appContext.GetRunningPort()))
}
