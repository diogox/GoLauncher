package extensions

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/models"
	"github.com/diogox/GoLauncher/extensions/websockets"
)

func StartExtensions(db *api.DB, responseChannel chan *api.Response) {

	// Start listening for extensions
	go websockets.StartExtensionsServer(responseChannel)

	// TODO: Run extensions in db (Receive db through args in this function)
	_ = (*db).AddExtension(models.Extension{
		Keyword: "hn",
		Name: "HackerNews",
		Description: "See the top stories on hackernews.",
		IconName: "google",
		DeveloperName: "Diogo Xavier",
	})
}