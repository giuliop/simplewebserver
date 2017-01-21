package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strings"
)

var port = flag.String("port", "8080", "port to use for http traffic (default is 8080)")
var ssl = flag.Bool("ssl", false, "whether you want to listen for HTTPS traffic (default is false)")
var ssl_port = flag.String("ssl_port", "8443", "port to use for https traffic (default is 8443)")
var ssl_cert = flag.String("ssl_cert", "ssl.crt", "file address of the ssl certificate (default is ssl.crt)")
var ssl_key = flag.String("ssl_key", "ssl.key", "file address of the ssl private key (default is ssl.key)")
var secure = flag.Bool("secure", true, "whether you want to redirect all http traffic to https (default is true)")

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	startMsg := "listening on port " + *port + " (http)"
	if *ssl {
		startMsg += " and on port " + *ssl_port + " (https)"
		go listenSSL(*ssl_port, *ssl_cert, *ssl_key)
	}
	log.Printf("%v...", startMsg)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func listenSSL(port, cert, key string) {
	err := http.ListenAndServeTLS(":"+port, cert, key, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// redirect all http traffic to https if 'secure' flag set
	if *ssl && *secure && r.TLS == nil {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusFound)
	} else {
		requester, err := net.LookupAddr(strings.Split(r.RemoteAddr, ":")[0])
		if err != nil {
			log.Print(err)
		}
		log.Printf("%v %v %v", requester, r.Method, r.URL)
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	}
}
