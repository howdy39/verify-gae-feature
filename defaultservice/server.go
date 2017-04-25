package backend

import (
	"net/http"
	appengineController "verify-gae-feature/src/controllers/appengine"
	logController "verify-gae-feature/src/controllers/log"
	rootController "verify-gae-feature/src/controllers/root"
)

func init() {
	http.HandleFunc("/", rootController.Root)
	http.HandleFunc("/appengine/", appengineController.AppenginePackage)
	http.HandleFunc("/log/", logController.Log)
}
