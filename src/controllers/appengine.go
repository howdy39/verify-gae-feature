package controllers

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// AppenginePackage appengine-packageのコントローラ
func AppenginePackage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	log.Debugf(c, "appID: %v", appengine.AppID(c))
	log.Debugf(c, "Datacenter: %v", appengine.Datacenter(c))
	log.Debugf(c, "DefaultVersionHostname: %v", appengine.DefaultVersionHostname(c))
	log.Debugf(c, "IsDevAppServer: %v", appengine.IsDevAppServer())
	log.Debugf(c, "VersionID: %v", appengine.VersionID(c))

	if err := appenginePackageTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var appenginePackageTemplate = template.Must(template.New("appengine").Parse(`
<html>
  <head>
    <title>verify gae feature</title>
  </head>
  <body>
    <h1>Reference</h1>
	<ul>
	  <li>
	    <a href="https://cloud.google.com/appengine/docs/standard/go/reference?hl=ja">https://cloud.google.com/appengine/docs/standard/go/reference?hl=ja</a>
	  </li>
	</ul>
  </body>
</html>
`))
