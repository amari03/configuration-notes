// Filename: main.go
package main

import (
	"flag"
    "log"
    "net/http"
)

// A handler function named 'home'
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from UBIT newsletter "))
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
    // retrieve the command line arguments
    flag.Parse()

    // mux is our router
    mux := http.NewServeMux()
    // the route pattern/endpoint/URL path
    mux.HandleFunc("/", home)

	log.Printf("starting server on %s", *addr)
    err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)


}

