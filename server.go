package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Connection established")
    })
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hi")

        //Display the method
        fmt.Fprintln(w, r.Method)

        //Display the path
        fmt.Fprintln(w, r.URL.Path)

        fmt.Fprintln(w, r.Header)

        fmt.Fprintln(w, http.StatusOK)

        fmt.Fprintln(w, "Proxy is ready")
    }) 
    port := ":5555"
    fmt.Println("Server is running on port" + port)

    log.Fatal(http.ListenAndServe(port, nil))
}