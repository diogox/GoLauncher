package events

import (
	"github.com/diogox/GoLauncher/api"
	"strings"
)

func KeywordQueryNew(query string) KeywordQuery {
	return KeywordQuery{
		query: query,
	}
}

type KeywordQuery struct {
	api.BaseEvent
	query string
}

func (kqe KeywordQuery) Keyword() string {
	return strings.Split(kqe.query, " ")[0]
}

func (kqe KeywordQuery) Argument() []string {
	return strings.Split(kqe.query, " ")[1:]
}
