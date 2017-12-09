package intent

import (
	"fmt"

	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/api"
	"github.com/train-cat/bot/helper"
	"github.com/train-cat/bot/wording"
)

const (
	ActionAskingForCreateAlert      = "asking_for_create_alert"
	ActionCreateAlert               = "create_alert"
	ActionCreateAlertSelectStoptime = "create_alert_select_stoptime"
	ActionCreateAlertRetry          = "create_alert_retry"

	ctxCreateAlert     = "ctx_create_alert"
	evtCreateAlert     = "event_create_alert"
	evtSelectStopsTime = "event_create_alert_select_stopstime"

	keyOriginID       = "origin_id"
	keyDestinationID  = "destination_id"
	keySchedule       = "schedule"
	keyChoiceStoptime = "choice_stoptime"
)

func forwardToCreateAlert(res *dialogflow.Response, say string) (*dialogflow.Response, error) {
	res.FollowUpEvent = &dialogflow.FollowUpEvent{
		Name: evtCreateAlert,
	}

	res.ContextOut = dialogflow.Contexts{
		{Name: ctxCreateAlert, Lifespan: 3, Parameters: dialogflow.Parameters{"say": say}},
		{Name: "askingforcreatealert-followup", Lifespan: 3},
	}

	return res, nil
}

func AskingForCreateAlert(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	return forwardToCreateAlert(res, wording.Get(wording.StartCreateAlert))
}

func CreateAlert(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	// request via event
	if req.Result.ResolvedQuery == evtCreateAlert {
		ctx, _ := req.Result.Contexts.Find(ctxCreateAlert)
		say, _ := ctx.Parameters.GetString("say")
		res.AddText(dialogflow.TextMessage{Speech: say}, helper.Platforms...)
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

	stopsTime, err := api.SearchStops(originID, destinationID, schedule)

	if err != nil {
		return helper.BotHasFail(res, err)
	}

	if len(stopsTime) == 0 {
		return forwardToCreateAlert(res, wording.Get(
			wording.NoStopTime, api.FindStationNameByID(originID), api.FindStationNameByID(destinationID), helper.FormatSchedule(schedule)),
		)
	}

	if req.Result.ResolvedQuery == evtSelectStopsTime {
		res.AddText(dialogflow.TextMessage{Speech: wording.Get(wording.AskListSchedule)}, helper.Platforms...)

		// TODO add oral version

		for i, stoptime := range stopsTime {
			str := fmt.Sprintf("Choix %d", i+1)

			res.AddCard(dialogflow.CardMessage{
				ImageURL: helper.StopTimeToCards(stoptime),
				Title:    str,
				// Subtitle: "Je suis sous-titre",
				Buttons: []dialogflow.Button{
					{Text: "Je choisis celui-là ☝️", PostBack: str},
				},
			}, helper.Platforms...)
		}

		return res, nil
	}

	if !req.Result.Parameters.HasKey(keyChoiceStoptime) {
		res.AddText(dialogflow.TextMessage{
			Speech: wording.Get(wording.ReAskSelectSchedule),
		}, helper.Platforms...)

		return res, nil
	}

	choice, _ := req.Result.Parameters.GetInt(keyChoiceStoptime)

	if choice > len(stopsTime) {
		res.AddText(dialogflow.TextMessage{
			Speech: wording.Get(wording.ChoiceOutOfRange, len(stopsTime)),
		}, helper.Platforms...)

		return res, nil
	}

	stoptime := stopsTime[choice-1]

	if err = api.CreateAlert(originID, stoptime.ID, req.OriginalRequest.Source, req.GetUserID()); err != nil {
		return helper.BotHasFail(res, err)
	}

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.ConfirmationAlert),
	}, helper.Platforms...)

	return res, nil
}

func CreateAlertRetry(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	return forwardToCreateAlert(res, wording.Get(wording.Retry))
}
