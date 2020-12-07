package main

import (
	"fmt"
	"os"

	"github.com/isurusiri/services/foo/context"
)

var app App
var appContext context.AppContext

func initContext() {
	appContext := context.New()
	appContext.SetRunningPort(os.Getenv("FOO_RUNNING_PORT"))
	appContext.SetRemoteHost(os.Getenv("FOO_REMOTE_HOST"))
}

func main() {
	fmt.Println("Starting foo api.")
	app.Initialize()
	initContext()

	app.Run(fmt.Sprintf(":%v", appContext.GetRunningPort()))
}
