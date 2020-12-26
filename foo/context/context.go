package context

// AppContext holds all configurations of the
// running application instance.
type AppContext struct {
	RunningPort string
	RemoteHost  string
}

var appContext *AppContext

// New create a new singleton instance of the app
func New(port string, remoteHost string) *AppContext {

	if appContext == nil {
		appContext = &AppContext{RunningPort: port, RemoteHost: remoteHost}
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

// GetRemoteHost returns the url of remote
func GetRemoteHost() string {
	return appContext.RemoteHost
}

// SetRemoteHost sets the remote port
func SetRemoteHost(remote string) {
	appContext.RemoteHost = remote
}
