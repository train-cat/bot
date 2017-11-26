package api

import (
	"github.com/spf13/viper"
	"github.com/train-cat/client-train-go"
)

// Init function, should be call after log & config init
func Init() {
	traincat.SetConfig(traincat.Config{
		Host: viper.GetString("api-train.host"),
		Auth: traincat.Auth{
			Username: viper.GetString("api-train.username"),
			Password: viper.GetString("api-train.password"),
		},
		Debug: viper.GetBool("bot.debug"),
	})
}
