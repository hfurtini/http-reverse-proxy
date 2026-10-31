package main

import(
	"log"
	"net/http"
)

func status(website string) (path string, status int){
    resp, err := http.Get(website)
	if err != nil{
		log.Fatalln(err)
	}

    return resp.Request.URL.Path ,resp.StatusCode
}