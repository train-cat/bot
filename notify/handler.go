package notify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/train-cat/bot/wording"
	"github.com/train-cat/client-train-go"
)

type (
	issue struct {
		Station  string `json:"station"`
		Schedule string `json:"schedule"`
		State    string `json:"state"`
	}

	action struct {
		Data map[string]string `json:"data"`
		Type string            `json:"type"`
	}

	notification struct {
		Issue   issue    `json:"issue"`
		Actions []action `json:"actions"`
	}
)

const (
	stateDelayed = "delayed"
	stateDeleted = "deleted"
)

// Handler notify request asking for notification
func Handler(w http.ResponseWriter, req *http.Request) {
	if !isTrustRequest(req) {
		log.Errorf("[handler notify] untrust request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	n := &notification{}

	bs, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Errorf("[handler notify] %s", err)
		return
	}

	err = json.Unmarshal(bs, n)

	if err != nil {
		log.Errorf("[handler notify] %s", err)
		return
	}

	go n.send()

	w.WriteHeader(http.StatusNoContent)
}

func isTrustRequest(req *http.Request) bool {
	return req.Method == http.MethodPost &&
		token == req.URL.Query().Get("token")
}

func (n notification) send() {
	var err error
	messageOne := wording.Get(wording.IssueOne, n.Issue.Station, n.Issue.Schedule, n.Issue.getFrenchState())
	messageTwo := wording.Get(wording.IssueTwo)

	for _, a := range n.Actions {
		switch a.Type {
		case traincat.ActionTypeTelegram:
			err = Telegram(a.getUserID(a.Type), messageOne, messageTwo)
		}

		if err != nil {
			log.Errorf("[send notification] %s", err)
		}
	}
}

func (i issue) getFrenchState() string {
	switch i.State {
	case stateDelayed:
		return "retardé"
	case stateDeleted:
		return "supprimé"
	default:
		return ""
	}
}

func (a action) getUserID(typ string) string {
	switch typ {
	case traincat.ActionTypeTelegram:
		userID, _ := a.Data["user_id"]

		return userID
	default:
		return ""
	}
}
