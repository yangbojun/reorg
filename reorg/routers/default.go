package routers

import (
	"net/http"
	"reorg/controllers"
)

func Register() {
	http.HandleFunc("/", controllers.WebServer)
}
