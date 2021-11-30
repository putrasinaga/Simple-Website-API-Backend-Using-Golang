package Handler

import (
	"html/template"
	"log"
	"path"

	"net/http"
)

func Input(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		templ, err := template.ParseFiles(path.Join("view", "form.html"))

		if err != nil {
			log.Print(err)
			http.Error(w, "parsing web error", 500)
			return

		}

		templ.Execute(w, nil)
		if err != nil {
			log.Print(err)
			http.Error(w, "gagal dijalankan", 500)
			return
		}
		return

	}

	http.Error(w, "method tidak di terima", 404)

}
