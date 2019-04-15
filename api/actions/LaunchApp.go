package actions

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/LinuxApps"
)

type LaunchAppOptions struct {
	Db             api.DB
	Exec           string
	SearchTermUsed string
}

func NewLaunchApp(opts LaunchAppOptions) LaunchApp {
	return LaunchApp{
		Type:           api.LAUNCH_APP_ACTION,
		Exec:           opts.Exec,
		db:             opts.Db,
		searchTermUsed: opts.SearchTermUsed,
	}
}

type LaunchApp struct {
	Type           string
	Exec           string
	db             api.DB
	searchTermUsed string
}

func (la LaunchApp) GetType() string {
	return la.Type
}

func (LaunchApp) KeepAppOpen() bool {
	return false
}

func (la LaunchApp) Run() error {
	// Increment access counter
	err := la.db.IncrementAppAccessCounter(la.Exec)
	if err != nil {

		// TODO: Log the error
		// If we can't increment the counter, we probably shouldn't return the error either, since it's not a critical functionality
		// TODO: I'll return the error for now, for testing purposes
		return err
	}

	err = LinuxApps.StartAppOrFocusExistingByCommand(la.Exec)
	if err != nil {
		return err
	}

	// Add session to db
	err = la.db.AddSession(la.searchTermUsed, la.Exec)
	if err != nil {

		// TODO: Log the error
		// If we can't increment the counter, we probably shouldn't return the error either, since it's not a critical functionality
		// TODO: I'll return the error for now, for testing purposes
		return err
	}

	return nil
}
