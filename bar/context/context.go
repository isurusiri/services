package context

// AppContext holds all configurations of the
// running application instance.
type AppContext struct {
	RunningPort string
}

var appContext *AppContext

// New create a new singleton instance of the app
func New(port string) *AppContext {

	if appContext == nil {
		appContext = &AppContext{RunningPort: port}
	}

	return appContext
}

// GetRunningPort returns the port
func GetRunningPort() string {
	return appContext.RunningPort
}

// SetRunningPort sets the running port
func SetRunningPort(port string) {
	appContext.RunningPort = port
}
