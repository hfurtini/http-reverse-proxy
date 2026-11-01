package main

import (
	"io"
	"log"
	"net/http"
)

func createRequest(website string, w http.ResponseWriter) (sb string){
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