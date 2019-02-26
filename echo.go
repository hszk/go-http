package main

import (
    "fmt"
    "net/http"
)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, r)
    }
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8888", nil)
}
