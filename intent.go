package main

import (
	"github.com/Eraac/dialogflow"
	"github.com/spf13/viper"
	"github.com/train-cat/bot/intent"
)

func getHandler() *dialogflow.Router {
	h := dialogflow.NewRouter(dialogflow.Config{
		Debug: viper.GetBool("bot.debug"),
		Token: viper.GetString("dialogflow.token"),
	})

	h.HandleFunc(intent.ActionWelcome, intent.Welcome)
	h.HandleFunc(intent.ActionCancel, intent.Cancel)
	h.HandleFunc(intent.ActionAskingForCreateAlert, intent.AskingForCreateAlert)
	h.HandleFunc(intent.ActionCreateAlert, intent.CreateAlert)
	h.HandleFunc(intent.ActionCreateAlertSelectStoptime, intent.CreateAlertSelectStoptime)

	return h
}
