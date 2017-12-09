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

var (
	// Platforms we have to interact with this bot
	Platforms = []string{dialogflow.PlatformDialogflow, dialogflow.PlatformTelegram, dialogflow.PlatformFacebook}
)
