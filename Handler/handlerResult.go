package Handler

import (
	"html/template"
	"log"
	"path"

	"net/http"
)

func Result(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "ada yang salah", 404)
			return
		}

		Firstname := r.Form.Get("firstname")
		Lastname := r.Form.Get("lastname")

		templ, err := template.ParseFiles(path.Join("view", "proses.html"))

		if err != nil {
			log.Print(err)
			http.Error(w, "parsing web proses error", 500)
			return

		}

		data := map[string]interface{}{
			"firstname": Firstname,
			"lastname":  Lastname,
		}

		templ.Execute(w, data)
		if err != nil {
			log.Print(err)
			http.Error(w, "gagal dijalankan", 500)
			return
		}
		return

	}

	http.Error(w, "method tidak ada", 404)

}
