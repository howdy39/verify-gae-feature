package datastore

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Datastore データストアの検証
func Datastore(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "DatastoreController")

	if err := rootTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<html>
  <head>
    <title>verify gae feature</title>
  </head>
  <body>
  	<ul>
	  <li>
	  	<a href="./create">create</a>
	  </li>
	  <li>
	  	<a href="./retrieve">retrieve</a>
	  </li>
	  <li>
	  	<a href="./update">update</a>
	  </li>
	  <li>
	  	<a href="./delete">delete</a>
	  </li>
	</ul>
  </body>
</html>
`))

// DatastoreCreate データストアの検証
func DatastoreCreate(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "DatastoreCreate")
}

// DatastoreRetrieve データストアの検証
func DatastoreRetrieve(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "DatastoreRetrieve")
}

// DatastoreUpdate データストアの検証
func DatastoreUpdate(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "DatastoreUpdate")
}

// DatastoreDelete データストアの検証
func DatastoreDelete(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "DatastoreDelete")
}
