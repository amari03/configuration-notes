package main

import (
    "net/http"
	"time"
)

func (app *application) loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            app.logger.Info("Method: %s | URI: %s | Time: %s",
                r.Method, r.URL.Path, start.Format(time.RFC3339))
            next.ServeHTTP(w, r)
        })
}