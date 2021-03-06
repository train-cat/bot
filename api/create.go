package api

import (
	"github.com/train-cat/client-train-go"
	"github.com/Eraac/dialogflow"
)

// CreateAlert -
func CreateAlert(stationID int, stopTimeID uint, source string, userID string) error {
	typ := sourceToActionType(source)

	actionInput := traincat.ActionInput{
		Type: typ,
		Data: generateData(typ, userID),
	}

	action, err := traincat.PostAction(actionInput)

	if err != nil {
		return err
	}

	alertInput := traincat.AlertInput{ActionID: action.ID}

	_, err = traincat.PostAlert(stationID, stopTimeID, alertInput)

	return err
}

func generateData(typ string, userID string) map[string]string {
	data := map[string]string{}

	switch typ {
	case traincat.ActionTypeMessenger:
		data["messenger_id"] = userID
	case traincat.ActionTypeTelegram:
		data["user_id"] = userID
	}

	return data
}

func sourceToActionType(source string) string {
	if source == dialogflow.PlatformFacebook {
		return traincat.ActionTypeMessenger
	}

	return source
}
