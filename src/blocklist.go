package main

import (
	"fmt"
	"log"
	"net/http"
)

func run_blocklist(w http.ResponseWriter, r *http.Request, website string){
	header := r.Header.Get("X-Block-Me")
    if header == "" || header == "false"{
        fmt.Fprintln(w, "Proxy is ready")
        fmt.Fprintln(w, createRequest(website))
    } else if header == "true"{
        fmt.Fprintln(w, http.StatusForbidden)
        log.Fatalln(w, " Blocked")
	}
}