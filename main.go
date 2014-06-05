package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "you can specify which port to use for http traffic (default is 8080)")
var ssl = flag.Bool("ssl", true, "you can specify if you want to listen for HTTPS traffic (default is true)")
var ssl_port = flag.String("ssl_port", "8443", "you can specify which port to use for https traffic (default is 8443)")
var secure = flag.Bool("secure", true, "you can specify if you want to redirect all http traffic to https (default is true)")

func main() {
	http.HandleFunc("/", handler)
	flag.Parse()
	startMsg := "listening on port " + *port + " (http)"
	if *ssl {
		startMsg += " and on port " + *ssl_port + " (https)"
		go func() {
			err := http.ListenAndServeTLS(":"+*ssl_port, "../ssl/giuliopizzini_com.crt", "../private/giuliopizzini.com.key.nopass", nil)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	log.Printf("%v...", startMsg)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// redirect all http traffic to https if 'secure' flag set
	if *secure && r.TLS == nil {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusFound)
	} else {
		log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL)
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	}
}
