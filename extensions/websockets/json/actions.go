package json

import (
	"encoding/json"
	"errors"
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/actions"
)

func InferActionType(jsonObj []byte) (*api.Action, error) {

	// Check if ActionList
	var actionList actions.ActionList
	err := json.Unmarshal(jsonObj, &actionList)
	if err == nil && actionList.Type == api.ACTION_LIST_ACTION {
		action := api.Action(actionList)
		return &action, nil
	}

	// Check if CopyToClipboard
	var copyToClipboard actions.CopyToClipboard
	err = json.Unmarshal(jsonObj, &copyToClipboard)
	if err == nil && copyToClipboard.Type == api.COPY_TO_CLIPBOARD_ACTION {
		action := api.Action(copyToClipboard)
		return &action, nil
	}

	// Check if DoNothing
	var doNothing actions.DoNothing
	err = json.Unmarshal(jsonObj, &doNothing)
	if err == nil && doNothing.Type == api.DO_NOTHING_ACTION {
		action := api.Action(doNothing)
		return &action, nil
	}

	// Check if HideWindow
	var hideWindow actions.HideWindow
	err = json.Unmarshal(jsonObj, &hideWindow)
	if err == nil && hideWindow.Type == api.HIDE_WINDOW_ACTION{
		action := api.Action(hideWindow)
		return &action, nil
	}

	// Check if LaunchApp
	var launchApp actions.LaunchApp
	err = json.Unmarshal(jsonObj, &launchApp)
	if err == nil && launchApp.Type == api.LAUNCH_APP_ACTION{
		action := api.Action(launchApp)
		return &action, nil
	}

	// Check if Open
	var open actions.Open
	err = json.Unmarshal(jsonObj, &open)
	if err == nil && open.Type == api.OPEN_ACTION{
		action := api.Action(open)
		return &action, nil
	}

	// Check if OpenUrl
	var openUrl actions.OpenUrl
	err = json.Unmarshal(jsonObj, &openUrl)
	if err == nil && openUrl.Type == api.OPEN_URL_ACTION{
		action := api.Action(openUrl)
		return &action, nil
	}

	// Check if RenderResultList
	var renderResultList actions.RenderResultList
	err = json.Unmarshal(jsonObj, &renderResultList)
	if err == nil && renderResultList.Type == api.RENDER_RESULT_LIST_ACTION{
		action := api.Action(renderResultList)
		return &action, nil
	}

	// Check if SetUserQuery
	var setUserQuery actions.SetUserQuery
	err = json.Unmarshal(jsonObj, &setUserQuery)
	if err == nil && setUserQuery.Type == api.SET_USER_QUERY_ACTION{
		action := api.Action(setUserQuery)
		return &action, nil
	}

	return nil, errors.New("action not recognized")
}
