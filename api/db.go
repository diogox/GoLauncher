package api

import "github.com/diogox/GoLauncher/api/models"

type DB interface {

	// Apps
	AddApp(app models.AppInfo) error
	IncrementAppAccessCounter(exec string) error
	UpdateApp(app models.AppInfo) error
	RemoveAllApps() error
	GetAllApps() ([]models.AppInfo, error)
	GetMostFrequentApps(nOfResults int) ([]models.AppInfo, error)
	FindAppByID(exec string) (models.AppInfo, error)
	FindAppByName(name string) ([]models.AppInfo, error)

	// Sessions
	AddSession(input string, appExec string) error
	GetSession(input string) (string, error)

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

	// Extensions
	AddExtension(extension models.Extension) error
	DeleteExtension(extension models.Extension) error
	GetAllExtensions() ([]models.Extension, error)
	FindExtensionByKeyword(keyword string) (models.Extension, error)
}