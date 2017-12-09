package intent

import (
	"github.com/Eraac/dialogflow"
	"github.com/train-cat/bot/wording"
)

// ActionHelp -
const ActionHelp = "help"

// Help intent
func Help(req *dialogflow.Request) (*dialogflow.Response, error) {
	res := dialogflow.NewResponse()

	res.AddQuickReply(dialogflow.QuickReplyMessage{
		Title: wording.Get(wording.Help),
		Replies: []string{
			wording.Get(wording.HelloTwoReplies),
		},
	}, req.Source())

	return res, nil
}
