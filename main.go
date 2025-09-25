package main

import (
    "log"
    "net/http"

    "example.com/go-ci-sample/handlers"
)

func main() {
    http.HandleFunc("/health", handlers.Health)
    http.HandleFunc("/add", handlers.AddHandler)

    addr := ":8080"
    log.Printf("starting server on %s", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatal(err)
    }
}
