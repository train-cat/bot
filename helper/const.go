package helper

// List of exit code available
const (
	ExitCodeSuccess = iota
	ExitCodeErrorInitConfig
	ExitCodeErrorListenServer
	ExitCodeErrorStopServer
	ExitCodeErrorInitNotification
)
