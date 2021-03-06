package api

type Event interface {
	GetEventType() string
}

const ITEM_ENTER_EVENT = "ITEM_ENTER_EVENT"
const KEYWORD_QUERY_EVENT = "KEYWORD_QUERY_EVENT"
const PREFERENCES_EVENT = "PREFERENCES_EVENT"
const PREFERENCES_UPDATE_EVENT = "PREFERENCES_UPDATE_EVENT"
const SYSTEM_EXIT_EVENT = "SYSTEM_EXIT_EVENT"