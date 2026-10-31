package main

import(
	"log"
	"net/http"
	"io"
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