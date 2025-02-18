// Filename: main.go
package main

import (
    "log"
    "net/http"
)

// A handler function named 'home'
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from UBIT newsletter "))
}

func main() {

    // mux is our router
    mux := http.NewServeMux()
    // the route pattern/endpoint/URL path
    mux.HandleFunc("/", home)

    log.Print("starting server on :4000")

    // start a local web server
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}

