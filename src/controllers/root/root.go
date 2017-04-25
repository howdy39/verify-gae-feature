package root

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Root /のコントローラ
func Root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "RootController")

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
	  	<a href="./appengine">appengine</a>
	  </li>
	  <li>
	  	<a href="./log">log</a>
	  </li>
	  <li>
	  	<a href="./datastore">datastore</a>
	  </li>
	</ul>
  </body>
</html>
`))
