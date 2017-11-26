package helper

import "github.com/Eraac/dialogflow"

const (
	ExitCodeSuccess = iota
	ExitCodeErrorInitConfig
	ExitCodeErrorListenServer
	ExitCodeErrorStopServer
	ExitCodeErrorInitNotification
)

var (
	Platforms = []string{dialogflow.PlatformDialogflow, dialogflow.PlatformTelegram}
)
