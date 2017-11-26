package api

import (
	"strings"

	"github.com/train-cat/client-train-go"
)

// GetCodeTrainFromStop return code train
func GetCodeTrainFromStop(s *traincat.Stop) string {
	ss := strings.Split(s.Links["train"].Href, "/")

	if len(ss) != 3 {
		return ""
	}

	return ss[2]
}
