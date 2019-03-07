package events

import (
	"github.com/diogox/GoLauncher/api"
	"strings"
)

func KeywordQueryNew(query string) KeywordQuery {
	return KeywordQuery{
		Type:  api.KEYWORD_QUERY_EVENT,
		Query: query,
	}
}

type KeywordQuery struct {
	Type  string
	Query string
}

func (kq KeywordQuery) GetEventType() string {
	return kq.Type
}

func (kq KeywordQuery) Keyword() string {
	return strings.Split(kq.Query, " ")[0]
}

func (kq KeywordQuery) Argument() []string {
	return strings.Split(kq.Query, " ")[1:]
}
