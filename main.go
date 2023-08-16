package main

import (
	"html/template"
	"net/http"
)

func main() {
	t := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/bread", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 Forbidden"))
			return
		}
		if r.Form.Has("type") {
			err = t.ExecuteTemplate(w, "bread/"+r.Form.Get("type"), nil)
			if err != nil {
				w.Write([]byte("Template not found!"))
			}
		}
	})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}
