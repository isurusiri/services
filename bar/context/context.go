package context

import "sync"

// AppContext holds all configurations of the
// running application instance.
type AppContext struct {
	runningPort string
}

var appContext *AppContext
var once sync.Once

// New create a new singleton instance of the app
func New() *AppContext {
	once.Do(func() {
		appContext = &AppContext{}
	})

	return appContext
}

// GetRunningPort returns the port
func (a *AppContext) GetRunningPort() string {
	return a.runningPort
}

// SetRunningPort sets the running port
func (a *AppContext) SetRunningPort(port string) {
	a.runningPort = port
}
