package notify

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/train-cat/bot/helper"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	telegram *tgbotapi.BotAPI
	token string
)

// Init should be call after log & config init
func Init() {
	var err error

	telegram, err = initTelegram()
	exitOnError(err)

	token = viper.GetString("bot.token")
}

func exitOnError(err error) {
	if err != nil {
		log.Errorf("[init notification] %s", err)
		os.Exit(helper.ExitCodeErrorInitNotification)
	}
}

func initTelegram() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(viper.GetString("bot.telegram.token"))

	if err != nil {
		return nil, err
	}

	bot.Debug = viper.GetBool("bot.debug")

	return bot, err
}
