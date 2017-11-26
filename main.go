package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/train-cat/bot/helper"
	"github.com/train-cat/bot/notify"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// === route ===
	http.HandleFunc("/notify", notify.Handler)
	http.Handle("/webhook", getHandler())
	http.HandleFunc("/health_check", healthCheck)
	http.Handle("/", http.NotFoundHandler())

	h := http.Server{
		Addr: fmt.Sprintf(":%d", viper.GetInt("http.port")),
	}

	go func() {
		log.Infof("HTTP Server listening on %s", h.Addr)

		if err := h.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
			os.Exit(helper.ExitCodeErrorListenServer)
		}
	}()

	<-stop

	log.Info("Graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.Shutdown(ctx)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(helper.ExitCodeErrorStopServer)
	}
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
