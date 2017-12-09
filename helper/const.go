package helper

import "github.com/Eraac/dialogflow"

// List of exit code available
const (
	ExitCodeSuccess = iota
	ExitCodeErrorInitConfig
	ExitCodeErrorListenServer
	ExitCodeErrorStopServer
	ExitCodeErrorInitNotification
)
