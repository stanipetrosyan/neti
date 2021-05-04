package main

import (
	"net/http"
	"text/template"
)

func main() {

	loginTmp := template.Must(template.ParseFiles("login.html"))
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		loginTmp.Execute(w, nil)
	})
	http.ListenAndServe(":8080", nil)
}
