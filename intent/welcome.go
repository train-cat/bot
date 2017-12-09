package intent

import (
	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/wording"
)

// ActionWelcome -
const ActionWelcome = "welcome"

// Welcome intent
func Welcome(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.HelloOne),
	}, req.Source())

	res.AddQuickReply(dialogflow.QuickReplyMessage{
		Title: wording.Get(wording.HelloTwo),
		Replies: []string{
			wording.Get(wording.HelloTwoReplies),
		},
	}, req.Source())

	res.ResetContext(req)

	return res, nil
}
