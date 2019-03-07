package websockets

import (
	"github.com/diogox/GoLauncher/api"
	"github.com/diogox/GoLauncher/api/models"
)

func StartExtensions(db *api.DB) {

	// Start listening for extensions
	go GetExtensionsServer().Start()

	// TODO: Run extensions in db (Receive db through args in this function)
	_ = (*db).AddExtension(models.Extension{
		Keyword: "hn",
		Name: "HackerNews",
		Description: "See the top stories on hackernews.",
		IconName: "google",
		DeveloperName: "Diogo Xavier",
	})
}