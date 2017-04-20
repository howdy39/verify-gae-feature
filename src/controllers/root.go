package controllers

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Root /のコントローラ
func Root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "appID: %v", appengine.AppID(c))

	if err := rooteTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var rooteTemplate = template.Must(template.New("root").Parse(`
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
	</ul>
  </body>
</html>
`))
