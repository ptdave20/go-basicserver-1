package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request URL: %s\r\n", r.RequestURI)
	fmt.Fprintf(w, "Method: %s\r\n", r.Method)
	fmt.Fprintf(w, "User Agent: %s\r\n", r.UserAgent())
	fmt.Fprintf(w, "Remote Address: %s\r\n", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}
