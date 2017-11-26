package intent

import (
	"strconv"

	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/api"
	"github.com/train-cat/bot/helper"
	"github.com/train-cat/bot/wording"
)

// List of actions receive by dialogflow
const (
	ActionCreateAlertStart      = "create_alert_start"
	ActionCreateAlert           = "create_alert"
	ActionCreateAlertSelectStop = "create_alert_select_stop"
	ActionCreateAlertConfirm    = "create_alert_confirm"
	ActionCreateAlertNo         = "create_alert_no"
)

const (
	keyOriginID       = "origin_id"
	keyOriginRaw      = "origin_raw"
	keyDestinationID  = "destination_id"
	keyDestinationRaw = "destination_raw"
	keySchedule       = "schedule"
	keyCode           = "code"
	keyListEnabled    = "list_enabled"

	ctxGeneric          = "generic"
	ctxStartCreateAlert = "startcreatealert-followup"
	ctxCreateAlert      = "createalert-followup"
)

// CreateAlertStart intent
func CreateAlertStart(_ *dialogflow.Request) (*dialogflow.Response, error) {
	r := dialogflow.NewResponse()

	r.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.StartCreateAlert),
	}, helper.Platforms...)

	return r, nil
}

// CreateAlert intent
func CreateAlert(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	if req.Result.ActionIncomplete {
		if !req.Result.Parameters.HasKey(keyOriginID) {
			err := setOrigin(req, res)

			return res, err
		}

		if !req.Result.Parameters.HasKey(keyDestinationID) {
			err := setDestination(req, res)

			return res, err
		}

		origin, _ := req.Result.Parameters.GetInt(keyOriginID)
		destination, _ := req.Result.Parameters.GetInt(keyDestinationID)

		if !req.Result.Parameters.HasKey(keySchedule) {
			res.AddText(dialogflow.TextMessage{
				Speech: wording.Get(wording.OriginOkAskSchedule, api.FindStationNameByID(destination), api.FindStationNameByID(origin)),
			}, helper.Platforms...)

			return res, nil
		}
	}

	origin, _ := req.Result.Parameters.GetInt(keyOriginID)
	destination, _ := req.Result.Parameters.GetInt(keyDestinationID)
	schedule, _ := req.Result.Parameters.GetString(keySchedule)

	stops, err := api.SearchStops(origin, destination, schedule)

	if err != nil {
		return nil, err
	}

	var replies []string

	for _, stop := range stops {
		replies = append(replies, stop.Schedule)
	}

	res.AddQuickReply(dialogflow.QuickReplyMessage{
		Title:   wording.Get(wording.AskListSchedule),
		Replies: replies,
	}, helper.Platforms...)

	return res, nil
}

// CreateAlertSelectStop intent
func CreateAlertSelectStop(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	ctx, err := req.Result.Contexts.Find(ctxCreateAlert)

	if err != nil {
		return nil, err
	}

	origin, err := ctx.Parameters.GetInt(keyOriginID)

	if err != nil {
		return nil, err
	}

	schedule, err := req.Result.Parameters.GetString(keySchedule)

	if err != nil {
		return nil, err
	}

	stop := api.FindStopByOriginAndSchedule(origin, schedule)

	if stop == nil {
		// TODO
	}

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.AskConfirmationAlert, schedule[0:len(schedule)-3], api.FindStationNameByID(origin)),
	}, helper.Platforms...)

	res.ContextOut = dialogflow.Contexts{
		{
			Name: ctxCreateAlert,
			Parameters: dialogflow.Parameters{
				keyOriginID: strconv.Itoa(origin),
				keyCode:     api.GetCodeTrainFromStop(stop),
			},
			Lifespan: 2,
		},
	}

	return res, nil
}

// CreateAlertConfirm intent
func CreateAlertConfirm(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	ctx, err := req.Result.Contexts.Find(ctxCreateAlert)

	if err != nil {
		return nil, err
	}

	code, err := ctx.Parameters.GetString(keyCode)

	if err != nil {
		return nil, err
	}

	originID, err := ctx.Parameters.GetInt(keyOriginID)

	if err != nil {
		return nil, err
	}

	userID := req.GetUserID()

	err = api.CreateAlert(originID, code, req.OriginalRequest.Source, userID)

	if err != nil {
		return nil, err
	}

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.ConfirmationAlert),
	}, helper.Platforms...)

	return res, nil
}

// CreateAlertNo intent
func CreateAlertNo(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.CancelAlert),
	}, helper.Platforms...)

	res.ResetContext(req)

	return res, nil
}

// === HELPER ===
func setOrigin(req *dialogflow.Request, res *dialogflow.Response) error {
	if !req.Result.Parameters.HasKey(keyOriginRaw) {
		res.AddText(dialogflow.TextMessage{
			Speech: wording.Get(wording.AskOrigin),
		}, helper.Platforms...)

		return nil
	}

	raw, _ := req.Result.Parameters.GetString(keyOriginRaw)
	ctx, _ := req.Result.Contexts.Find(ctxGeneric)

	if ctx.Parameters.HasKey(keyListEnabled) {
		raw = req.Result.ResolvedQuery
		ctx.Parameters["list_enable"] = ""
	}

	stations, err := api.SearchStations(raw)

	if err != nil {
		return err
	}

	switch len(stations) {
	case 0:
		res.AddText(dialogflow.TextMessage{
			Speech: wording.Get(wording.OriginNotFound),
		}, helper.Platforms...)

	case 1:
		res.FollowUpEvent = &dialogflow.FollowUpEvent{
			Name: ctxStartCreateAlert,
			Data: dialogflow.Parameters{
				keyOriginID: stations[0].ID,
			},
		}
	default:
		var replies []string

		for _, station := range stations {
			replies = append(replies, station.Name)
		}

		res.AddQuickReply(dialogflow.QuickReplyMessage{
			Title:   wording.Get(wording.SelectOrigin),
			Replies: replies,
		}, helper.Platforms...)

		ctx.Parameters[keyListEnabled] = "1"
	}

	res.ContextOut = append(res.ContextOut, *ctx)

	return nil
}

func setDestination(req *dialogflow.Request, res *dialogflow.Response) error {
	origin, _ := req.Result.Parameters.GetInt(keyOriginID)

	if !req.Result.Parameters.HasKey(keyDestinationRaw) {
		res.AddText(dialogflow.TextMessage{
			Speech: wording.Get(wording.AskDestination, api.FindStationNameByID(origin)),
		}, helper.Platforms...)

		return nil
	}

	raw, _ := req.Result.Parameters.GetString(keyDestinationRaw)
	ctx, _ := req.Result.Contexts.Find(ctxGeneric)

	if ctx.Parameters.HasKey(keyListEnabled) {
		raw = req.Result.ResolvedQuery
		ctx.Parameters["list_enable"] = ""
	}

	stations, err := api.SearchStations(raw)

	if err != nil {
		return err
	}

	switch len(stations) {
	case 0:
		res.AddText(dialogflow.TextMessage{
			Speech: wording.Get(wording.DestinationNotFound),
		}, helper.Platforms...)

	case 1:
		res.FollowUpEvent = &dialogflow.FollowUpEvent{
			Name: ctxStartCreateAlert,
			Data: dialogflow.Parameters{
				keyDestinationID: stations[0].ID,
			},
		}
	default:
		var replies []string

		for _, station := range stations {
			replies = append(replies, station.Name)
		}

		res.AddQuickReply(dialogflow.QuickReplyMessage{
			Title:   wording.Get(wording.SelectDestination),
			Replies: replies,
		}, helper.Platforms...)

		ctx.Parameters[keyListEnabled] = "1"
	}

	res.ContextOut = append(res.ContextOut, *ctx)

	return nil
}
