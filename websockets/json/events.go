package json

import (
	"encoding/json"
	"errors"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/events"
)

func InferEventType(jsonObj []byte) (*api.Event, error) {

	// Check if Preferences
	var preferences events.Preferences
	err := json.Unmarshal(jsonObj, &preferences)
	if err == nil && preferences.Type == api.PREFERENCES_EVENT {
		event := api.Event(preferences)
		return &event, nil
	}

	// Check if PreferencesUpdate
	var preferencesUpdate events.PreferencesUpdate
	err = json.Unmarshal(jsonObj, &preferencesUpdate)
	if err == nil && preferencesUpdate.Type == api.PREFERENCES_UPDATE_EVENT {
		event := api.Event(preferencesUpdate)
		return &event, nil
	}

	// Check if ItemEnter
	var itemEnter events.ItemEnter
	err = json.Unmarshal(jsonObj, &itemEnter)
	if err == nil && itemEnter.Type == api.ITEM_ENTER_EVENT {
		event := api.Event(itemEnter)
		return &event, nil
	}

	// Check if KeywordQuery
	var keywordQuery events.KeywordQuery
	err = json.Unmarshal(jsonObj, &keywordQuery)
	if err == nil && keywordQuery.Type == api.KEYWORD_QUERY_EVENT {
		event := api.Event(keywordQuery)
		return &event, nil
	}

	// Check if SystemExit
	var systemExit events.SystemExit
	err = json.Unmarshal(jsonObj, &systemExit)
	if err == nil && systemExit.Type == api.SYSTEM_EXIT_EVENT {
		event := api.Event(systemExit)
		return &event, nil
	}

	return nil, errors.New("event not recognized")
}
