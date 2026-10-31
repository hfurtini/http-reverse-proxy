package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func createRequest(website string) (sb string){
	resp, err := http.Get(website)
	if err != nil{
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
	}

    return string(body)
}

func status(website string) (path string, status int){
    resp, err := http.Get(website)
	if err != nil{
		log.Fatalln(err)
	}

    return resp.Request.URL.Path ,resp.StatusCode
}

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

        header := r.Header.Get("X-Block-Me")
        if header == ""{
            fmt.Fprintln(w, "Proxy is ready")
            fmt.Fprintln(w, createRequest("https://archlinux.org/"))
        } else if header == "true"{
            fmt.Fprintln(w, http.StatusForbidden)
            log.Fatalln(w, " Blocked")
        } else{
            fmt.Fprintln(w, "Proxy is ready")
            fmt.Fprintln(w, createRequest("https://archlinux.org/"))
        }

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

        header := r.Header.Get("X-Block-Me")
        if header == ""{
            fmt.Fprintln(w, "Proxy is ready")
            fmt.Fprintln(w, createRequest("https://httpbin.org/"))
        } else if header == "true"{
            fmt.Fprintln(w, http.StatusForbidden)
            fmt.Fprintln(w, "Blocked")
        } else{
            fmt.Fprintln(w, "Proxy is ready")
            fmt.Fprintln(w, createRequest("https://httpbin.org/"))
        }

        path, status := status("https://httpbin.org/")
        fmt.Fprintf(w, "[%s] [%s] - Status: [%d]", r.Method, path, status)
    }) 

    port := ":5555"
    fmt.Println("Server is running on port" + port)

    log.Fatal(http.ListenAndServe(port, nil))
}