package notify

import (
	"os"

	"github.com/abhinavdahiya/go-messenger-bot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/train-cat/bot/helper"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	telegram  *tgbotapi.BotAPI
	messenger *mbotapi.BotAPI
	token     string
)

var l logrus.FieldLogger

// Init should be call after log & config init
func Init(fl logrus.FieldLogger) {
	l = fl.WithField("action", "notify")

	var err error
	debug := viper.GetBool("bot.debug")

	telegram, err = initTelegram(debug)
	exitOnError(err)

	messenger = initMessenger(debug)

	token = viper.GetString("bot.token")
}

func exitOnError(err error) {
	if err != nil {
		l.WithField("status", "error").Error(err.Error())
		os.Exit(helper.ExitCodeErrorInitNotification)
	}
}

func initTelegram(debug bool) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(viper.GetString("bot.telegram.token"))

	if err != nil {
		return nil, err
	}

	bot.Debug = debug

	return bot, err
}

func initMessenger(debug bool) *mbotapi.BotAPI {
	bot := mbotapi.NewBotAPI(
		viper.GetString("bot.messenger.page_token"),
		viper.GetString("bot.messenger.verify_token"),
		viper.GetString("bot.messenger.secret_key"),
	)

	bot.Debug = debug

	return bot
}
