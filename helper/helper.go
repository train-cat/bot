package helper

import (
	"fmt"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/train-cat/client-train-go"
	"github.com/spf13/viper"
	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/wording"
)

// Int return pointer to int
func Int(i int) *int {
	return &i
}

// String return pointer to string
func String(s string) *string {
	return &s
}

// FormatSchedule transform 04:00:00 (from dialogflow) to 04:00
func FormatSchedule(schedule string) string {
	if len(schedule) >= 8 {
		return schedule[0:5]
	}

	return schedule
}

// StopTimeToCards return url of the image
func StopTimeToCards(s traincat.StopTime) string {
	var mission string
	var days int
	origin, _ := s.GetStation()
	terminus, _ := s.GetTerminus()

	s.Embedded.Get("mission", &mission)
	s.Embedded.Get("days", &days)

	q := url.Values{}
	q.Add("mission", mission)
	q.Add("days", strconv.Itoa(days))
	q.Add("schedule", s.Schedule)
	q.Add("origin", origin.Name)
	q.Add("terminus", terminus.Name)

	return fmt.Sprintf("%s/generate/stop_time?%s", viper.GetString("cards-generator.host"), q.Encode())
}

func BotHasFail(res *dialogflow.Response, err error) (*dialogflow.Response, error) {
	log.Error(err)

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.HasFail),
	}, Platforms...)

	return res, nil
}
