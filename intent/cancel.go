package intent

import (
	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/wording"
	"github.com/train-cat/bot/helper"
)

// ActionCancel -
const ActionCancel = "cancel"

// Cancel intent
func Cancel(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.Cancel),
	}, helper.Platforms...)

	res.ResetContext(req)

	return res, nil
}
