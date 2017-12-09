package notify

import (
	"strconv"

	"github.com/abhinavdahiya/go-messenger-bot"
)

// Messenger send message via Messenger API
func Messenger(userID string, messages ...string) error {
	id, err := strconv.Atoi(userID)

	if err != nil {
		return err
	}

	user := mbotapi.NewUserFromID(int64(id))

	for _, m := range messages {
		msg := mbotapi.NewMessage(m)

		_, err = messenger.Send(user, msg, mbotapi.RegularNotif)

		if err != nil {
			return err
		}
	}

	return nil
}
