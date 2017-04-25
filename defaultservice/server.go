package backend

import (
	"net/http"
	appengineController "verify-gae-feature/src/controllers/appengine"
	datastoreController "verify-gae-feature/src/controllers/datastore"
	logController "verify-gae-feature/src/controllers/log"
	rootController "verify-gae-feature/src/controllers/root"
)

func init() {
	http.HandleFunc("/", rootController.Root)
	http.HandleFunc("/appengine/", appengineController.AppenginePackage)
	http.HandleFunc("/log/", logController.Log)
	http.HandleFunc("/datastore/", datastoreController.Datastore)
	http.HandleFunc("/datastore/create/", datastoreController.DatastoreCreate)
	http.HandleFunc("/datastore/retrieve/", datastoreController.DatastoreRetrieve)
	http.HandleFunc("/datastore/update/", datastoreController.DatastoreUpdate)
	http.HandleFunc("/datastore/delete/", datastoreController.DatastoreDelete)
}
