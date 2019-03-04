package api

import "github.com/diogox/GoLauncher/api/models"

type DB interface {

	// Apps
	AddApp(exec string, name string, descr string, iconName string) error
	FindAppByID(exec string) (string, error)
	FindAppByName(name string) ([]models.AppInfo, error)

	// Preferences
	LoadDefaultPreferences() []error
	AddPreference(preference string, value string) error
	SetPreference(preference string, value string) error
	GetPreference(preference string) string

	// Shortcuts
	AddShortcut(shortcut models.ShortcutInfo) error
	DeleteShortcut(shortcut models.ShortcutInfo) error
	GetAllShortcuts() ([]models.ShortcutInfo, error)
	FindShortcutByKeyword(keyword string) (models.ShortcutInfo, error)
}