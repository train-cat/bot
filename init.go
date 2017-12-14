package main

import (
	"flag"
	"os"

	"github.com/Abramovic/logrus_influxdb"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/train-cat/bot/api"
	"github.com/train-cat/bot/helper"
	"github.com/train-cat/bot/notify"
)

func init() {
	initConfig()
	initLogging()
	api.Init()
	notify.Init(logrus.StandardLogger())
}

func initLogging() {
	config := &logrus_influxdb.Config{
		Database:    viper.GetString("influxdb.database"),
		Measurement: viper.GetString("influxdb.measurement"),
		Tags:        []string{"action", "intent", "status", "source", "session_id", "user_id", "type"},
	}

	// Connect to InfluxDB using the standard client.
	influxClient, _ := client.NewHTTPClient(client.HTTPConfig{
		Addr:     viper.GetString("influxdb.host"),
		Username: viper.GetString("influxdb.username"),
		Password: viper.GetString("influxdb.password"),
	})

	hook, err := logrus_influxdb.NewInfluxDB(config, influxClient)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.AddHook(hook)
}

func initConfig() {
	cfgFile := flag.String("config", "config.json", "config file")
	flag.Parse()

	viper.SetConfigFile(*cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		logrus.Error(err.Error())
		os.Exit(helper.ExitCodeErrorInitConfig)
	}
}
