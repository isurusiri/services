package main

import (
	"fmt"
	"os"

	"github.com/isurusiri/services/foo/context"
)

var app App
var appContext context.AppContext

func initContext() {
	appContext := context.New(os.Getenv("FOO_RUNNING_PORT"), os.Getenv("FOO_REMOTE_HOST"))
	fmt.Printf("AppContext is:%v\n", appContext)
}

func main() {
	fmt.Println("Starting foo api.")
	app.Initialize()
	initContext()

	app.Run(fmt.Sprintf(":%v", context.GetRunningPort()))
}
