//fserve is straight from the go example
//It's packaged in httputil/fserve so I can 'go get' it.
//Don't judge.
package main

import (
	"log"
	"net/http"
	"flag"
)

func main() {
	dir := flag.String("dir", ".", "Directory to serve")
	port := flag.String("port", ":8080", "Port to listen on.")
	flag.Parse()
	// Simple static webserver:
	log.Printf("Serving: %v", *dir)
	log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir(*dir))))
}
