package backend2

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", Root)
}

// Root /のコントローラ
func Root(w http.ResponseWriter, r *http.Request) {

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
    これはservice2です。
  </body>
</html>
`))
