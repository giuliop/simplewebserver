package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("p", "8080", "you can specify which port to use (default is 8080)")

func main() {
	http.HandleFunc("/", handler)
	flag.Parse()
	log.Printf("listening on port %v...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL)
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}
