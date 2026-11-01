package main

import(
	"net/http"
)

func preserveRealIp(w http.ResponseWriter, r *http.Request) string{
	clientIp := r.Header.Get("X-FORWARDED-FOR")
	if clientIp != ""{
		r.Header.Set("X-FORWARDED-FOR", clientIp)
		return clientIp
	} else{
		r.Header.Set("X-FORWARDED-FOR", r.RemoteAddr)
		return r.RemoteAddr
	}
}