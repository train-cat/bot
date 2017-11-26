package intent

import (
	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/wording"
)

const ActionCancel = "cancel"

func Cancel(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.Cancel),
	}, dialogflow.PlatformTelegram)

	res.ResetContext(req)

	return res, nil
}
