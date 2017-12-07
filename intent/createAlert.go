package intent

import (
	"fmt"

	"github.com/Eraac/dialogflow"
	// log "github.com/sirupsen/logrus"
	"github.com/train-cat/bot/api"
	"github.com/train-cat/bot/helper"
	"github.com/train-cat/bot/wording"
)

const (
	ActionAskingForCreateAlert      = "asking_for_create_alert"
	ActionCreateAlert               = "create_alert"
	ActionCreateAlertSelectStoptime = "create_alert_select_stoptime"

	ctxCreateAlert     = "ctx_create_alert"
	evtCreateAlert     = "event_create_alert"
	evtSelectStopsTime = "event_create_alert_select_stopstime"

	keyOriginID       = "origin_id"
	keyDestinationID  = "destination_id"
	keySchedule       = "schedule"
	keyChoiceStoptime = "choice_stoptime"
)

func AskingForCreateAlert(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.FollowUpEvent = &dialogflow.FollowUpEvent{
		Name: evtCreateAlert,
	}

	return res, nil
}

func CreateAlert(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	// requête via l'event de l'intent précédent
	if req.Result.ResolvedQuery == evtCreateAlert {
		res.AddText(dialogflow.TextMessage{Speech: wording.Get(wording.StartCreateAlert)}, helper.Platforms...)

		res.ContextOut = dialogflow.Contexts{
			{Name: ctxCreateAlert, Lifespan: 3},
		}
	}

	if !req.Result.Parameters.HasKey(keyOriginID) {
		res.AddText(dialogflow.TextMessage{Speech: wording.Get(wording.AskOrigin)}, helper.Platforms...)

		return res, nil
	}

	originID, _ := req.Result.Parameters.GetInt(keyOriginID)
	originName := api.FindStationNameByID(originID)

	if !req.Result.Parameters.HasKey(keyDestinationID) {
		res.AddText(dialogflow.TextMessage{Speech: wording.Get(wording.AskDestination, originName)}, helper.Platforms...)

		return res, nil
	}

	destinationID, _ := req.Result.Parameters.GetInt(keyDestinationID)
	destinationName := api.FindStationNameByID(destinationID)

	if !req.Result.Parameters.HasKey(keySchedule) {
		res.AddText(dialogflow.TextMessage{Speech: wording.Get(wording.AskSchedule, originName, destinationName)}, helper.Platforms...)

		return res, nil
	}

	res.FollowUpEvent = &dialogflow.FollowUpEvent{
		Name: evtSelectStopsTime,
	}

	return res, nil
}

func CreateAlertSelectStoptime(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	ctx, _ := req.Result.Contexts.Find(ctxCreateAlert)

	originID, _ := ctx.Parameters.GetInt(keyOriginID)
	destinationID, _ := ctx.Parameters.GetInt(keyDestinationID)
	schedule, _ := ctx.Parameters.GetString(keySchedule)

	stopsTime, _ := api.SearchStops(originID, destinationID, schedule)

	if req.Result.ResolvedQuery == evtSelectStopsTime {
		res.AddText(dialogflow.TextMessage{Speech: wording.Get(wording.AskListSchedule)}, helper.Platforms...)

		// TODO prévoir une version orale aussi

		for i, stoptime := range stopsTime {
			str := fmt.Sprintf("Choix %d", i+1)

			res.AddCard(dialogflow.CardMessage{
				ImageURL: fmt.Sprintf("https://cards-generator.train.cat/generate?mission=MALA&origin=Poissy&terminus=Paris%%20Saint-Lazare&schedule=%s&days=124", stoptime.Schedule),
				Title:    str,
				Subtitle: "Je suis sous-titre",
				Buttons: []dialogflow.Button{
					{Text: "Je choisis celui-là", PostBack: str},
				},
			}, helper.Platforms...)
		}

		return res, nil
	}

	if !req.Result.Parameters.HasKey(keyChoiceStoptime) {
		// TODO re-demander
	}

	choice, _ := req.Result.Parameters.GetInt(keyChoiceStoptime)

	if choice >= len(stopsTime) {
		// TODO ask user good choice
	}

	stoptime := stopsTime[choice -1]

	res.AddText(dialogflow.TextMessage{
		Speech: fmt.Sprintf("YES %d - %d", choice, stoptime.ID),
	}, helper.Platforms...)

	// TODO persist alert

	return res, nil
}
