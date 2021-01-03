package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"wiki/routes"
)

func main() {
	port:= flag.Int("port",3000,"port")
	flag.Parse()
	log.Printf("Starting at %d\n", *port)

	http.HandleFunc("/wiki/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			routes.HandleGetWiki(w,r)
		case "POST":
			routes.HandlePostWiki(w,r)
		default:
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintf(w, "not acceptable")
		}
	})

	http.ListenAndServe(fmt.Sprint(":",*port), nil)
}