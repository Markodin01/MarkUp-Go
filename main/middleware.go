package main

import (
    "fmt"
    "net/http"
)

// LoggingMiddleware is a simple middleware function that logs the request method and URL.
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Request method: %s, URL: %s\n", r.Method, r.URL)
        next.ServeHTTP(w, r)
    })
}

func main() {
    http.Handle("/", LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, World!")
    })))

    http.ListenAndServe(":8080", nil)
}