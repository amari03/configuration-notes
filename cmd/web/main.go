// Filename: main.go
package main

import (
	"flag"
	"log/slog"
   // "log"
    "net/http"
	"os"
)

// A handler function named 'home'
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from UBIT newsletter "))
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
    // retrieve the command line arguments
    flag.Parse()

	// Create a new structured logger
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // mux is our router
    mux := http.NewServeMux()
    // the route pattern/endpoint/URL path
    mux.HandleFunc("/", home)


	// Use the logger. Note the key:value pair
    logger.Info("starting server ", "addr", *addr)
    err := http.ListenAndServe(*addr, mux)
	//log.Printf("starting server on %s", *addr)
	// Use the logger. Note the error level
    logger.Error(err.Error())
    os.Exit(1)



}

