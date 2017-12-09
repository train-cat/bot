package intent

import (
	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/wording"
)

// ActionHelp -
const ActionThankYou = "thank_you"

// ThankYou intent
func ThankYou(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.AddText(dialogflow.TextMessage{Speech: wording.Get(wording.ThankYou)}, req.Source())

	return res, nil
}
