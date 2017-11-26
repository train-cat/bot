package api

import (
	"strings"

	"github.com/train-cat/client-train-go"
)

func GetCodeTrainFromStop(s *traincat.Stop) string {
	ss := strings.Split(s.Links["train"].Href, "/")

	if len(ss) != 3 {
		return ""
	}

	return ss[2]
}
