package controllers

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Log ログの検証
func Log(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	log.Debugf(c, "これはDebugログ")
	log.Warningf(c, "これはWarningログ")
	log.Infof(c, "これはInfoログ")
	log.Errorf(c, "これはErrorログ")
	log.Criticalf(c, "これはCriticalログ")
}
