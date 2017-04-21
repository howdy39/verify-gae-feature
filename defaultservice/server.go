package backend

import (
	"net/http"
	"verify-gae-feature/src/controllers"
)

func init() {
	http.HandleFunc("/", controllers.Root)
	http.HandleFunc("/appengine", controllers.AppenginePackage)
	http.HandleFunc("/log", controllers.Log)

}
