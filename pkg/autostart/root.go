package autostart

import (
	"github.com/ProtonMail/go-autostart"
)

func SetAppStart(isSet bool) {
	app := &autostart.App{
		Name: "GoLauncher", // TODO: Get this from config file?
		Exec:  []string{"go", "run", "/home/diogox/go/src/github.com/diogox/GoLauncher"}, // TODO: Change this to actual executable name when ready for distribution!!
		DisplayName: "GoLauncher",
		Icon: "google", // TODO: Change this to golauncher's logo
	}
	if app.IsEnabled() && !isSet {
		if err := app.Disable(); err != nil {
			panic(err)
		}
	} else if !app.IsEnabled() && isSet {
		if err := app.Enable(); err != nil {
			panic(err)
		}
	}
}
