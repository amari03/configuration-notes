// Filename: main.go
package main

import (
	"log/slog"
    "net/http"
	"os"
)

//dependency injection
type application struct{
    logger *slog.Logger
}

// A handler function named 'home'
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from UBIT newsletter "))
}

func main() {

    // mux is our router
    mux := http.NewServeMux()
    // the route pattern/endpoint/URL path
    mux.HandleFunc("/", home)
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    //create an instance of application
    app := &application{
        logger: logger,
    }
    app.logger.Info("starting server :4000")

    logger.Info("starting server :4000")
    err := http.ListenAndServe(":4000", app.loggingMiddleware(mux))
    // Use the logger. Note the error level
    logger.Error(err.Error())


}

