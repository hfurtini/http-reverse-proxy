package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

    http.HandleFunc("/server2", func(w http.ResponseWriter, r *http.Request) {
        //Display the method
        fmt.Fprintln(w, r.Method)

        //Display the path
        fmt.Fprintln(w, r.URL.Path)

        //Display the header
        fmt.Fprintln(w, r.Header)

        //Display the status
        fmt.Fprintln(w, http.StatusOK)

        run_blocklist(w, r, "https://archlinux.org/")

        path, status := status("https://archlinux.org/")
        fmt.Fprintf(w, "[%s] [%s] - Status: [%d]", r.Method, path, status)

    })
    http.HandleFunc("/server1", func(w http.ResponseWriter, r *http.Request) {

        //Display the method
        fmt.Fprintln(w, r.Method)

        //Display the path
        fmt.Fprintln(w, r.URL.Path)

        //Display the header
        fmt.Fprintln(w, r.Header)

        //Display the status
        fmt.Fprintln(w, http.StatusOK)

        run_blocklist(w, r, "https://httpbin.org/")

        path, status := status("https://httpbin.org/")
        fmt.Fprintf(w, "[%s] [%s] - Status: [%d]", r.Method, path, status)
    }) 

    port := ":5555"
    fmt.Println("Server is running on port" + port)

    log.Fatal(http.ListenAndServe(port, nil))
}