package main

import (
	"flag"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

var (
	port = flag.Int("port", 8080, "Port number.")

	statusCode = [...]int{
		200, 200, 200, 200, 200, 200, 200, 200, 200, 200,
		300, 301, 302, 304, 307,
		400, 401, 403, 403, 404, 404, 404, 404, 404, 410,
		500, 500, 500, 501, 503, 503, 503, 503, 503, 550}
)

func index(w http.ResponseWriter, r *http.Request) {
	type HTTPStatus struct {
		Code int
	}

	randomStatus := statusCode[rand.Intn(len(statusCode))]
	status := HTTPStatus{Code: randomStatus}
	w.WriteHeader(randomStatus)

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	tpl.ExecuteTemplate(w, "index.gohtml", status)
}

func main() {
	fmt.Printf("Listening port 8080\n")

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	flag.Parse()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
