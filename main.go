package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

var port = flag.String("p", "8080", "you can specify which port to use (default is 8080)")

func main() {
	http.HandleFunc("/", HomeHandler)
	log.Printf("listening on port %v...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
	if !strings.Contains(r.URL.Path, ".") {
		r.URL.Path = "/"
	}
	http.FileServer(http.Dir("./")).ServeHTTP(w, r)
}
