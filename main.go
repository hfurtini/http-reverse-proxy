package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func createRequest() (sb string){
	resp, err := http.Get("https://httpbin.org/")
	if err != nil{
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
	}

    return string(body)
}

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Connection established")
    })
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {

        //Display the method
        fmt.Fprintln(w, r.Method)

        //Display the path
        fmt.Fprintln(w, r.URL.Path)

        //Display the header
        fmt.Fprintln(w, r.Header)

        //Display the status
        fmt.Fprintln(w, http.StatusOK)

        fmt.Fprintln(w, "Proxy is ready")

        fmt.Fprintln(w, createRequest())
    }) 
    port := ":5555"
    fmt.Println("Server is running on port" + port)

    log.Fatal(http.ListenAndServe(port, nil))
}