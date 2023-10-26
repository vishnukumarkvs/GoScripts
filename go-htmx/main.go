package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	//handler
    h1 := func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w,"Hello World\n")
		io.WriteString(w, r.Method)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000",nil))
}