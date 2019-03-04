package models

type ShortcutInfo struct {
	Keyword         string
	Name            string
	Cmd             string
	IconName        string
	IsDefaultSearch bool
	IsActive        bool
}
