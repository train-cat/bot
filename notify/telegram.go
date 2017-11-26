package notify

import "net/url"

// Telegram send message via telegram API
func Telegram(userID string, messages ...string) error {
	for _, m := range messages {
		params := url.Values{}

		params.Add("chat_id", userID)
		params.Add("text", m)

		_, err := telegram.MakeRequest("sendMessage", params)

		if err != nil {
			return err
		}
	}

	return nil
}
