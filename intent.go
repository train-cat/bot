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
	h.HandleFunc(intent.ActionCreateAlertStart, intent.CreateAlertStart)
	h.HandleFunc(intent.ActionCreateAlert, intent.CreateAlert)
	h.HandleFunc(intent.ActionCreateAlertSelectStop, intent.CreateAlertSelectStop)
	h.HandleFunc(intent.ActionCreateAlertConfirm, intent.CreateAlertConfirm)
	h.HandleFunc(intent.ActionCreateAlertNo, intent.CreateAlertNo)

	return h
}
