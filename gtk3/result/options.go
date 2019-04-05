package result

import "github.com/diogox/GoLauncher/api"

type ResultItemOptions struct {
	Title            string
	Description      string
	IconPath         string
	IsDefaultSelect  bool
	OnEnterAction    api.Action
	OnAltEnterAction api.Action
}
