package extensions

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/extensions/websockets"
)

func StartExtensions(responseChannel chan *api.Response) {

	// Start listening for extensions
	go websockets.StartExtensionsServer(responseChannel)

	// TODO: Run extensions in db (Receive db through args in this function)
}