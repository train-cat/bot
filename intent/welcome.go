package intent

import (
	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/wording"
	"github.com/train-cat/bot/helper"
)

// ActionWelcome -
const ActionWelcome = "welcome"

// Welcome intent
func Welcome(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.AddText(dialogflow.TextMessage{
		Speech: wording.Get(wording.HelloOne),
	}, helper.Platforms...)

	res.AddQuickReply(dialogflow.QuickReplyMessage{
		Title: wording.Get(wording.HelloTwo),
		Replies: []string{
			wording.Get(wording.HelloTwoReplies),
		},
	}, helper.Platforms...)

	res.ResetContext(req)

	return res, nil
}
