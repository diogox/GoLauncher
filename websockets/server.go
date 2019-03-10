package websockets

import (
	"github.com/diogox/GoLauncher/api/models"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var instance *ExtensionsServer
var once sync.Once

func GetExtensionsServer() *ExtensionsServer {
	once.Do(func() {
		instance = &ExtensionsServer{}
	})
	return instance
}

type ExtensionsServer struct {
	controllers []*Controller
}

func (es *ExtensionsServer) Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		// TODO: Get actual extension info
		extension := models.Extension{
			Keyword:       "hn",
			Name:          "HackerNews",
			Description:   "See the top stories on HackerNews!",
			IconName:      "google",
			DeveloperName: "Diogo Xavier",
		}
		controller := ControllerNew(&extension, conn)
		es.controllers = append(es.controllers, controller)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

func (es *ExtensionsServer) GetControllerByKeyword(keyword string) *Controller {
	for _, controller := range es.controllers {
		if controller.Extension.Keyword == keyword {
			return controller
		}
	}
	return nil
}