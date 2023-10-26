package main

import (
	"log"
	"net/http"
	"text/template"
)

type Film struct{
	Title string
	Director string
}

func main() {

	//handler
    h1 := func(w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("./index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "Salaar", Director: "Prashant Neel"},
				{Title: "Pushpa", Director: "Sukumar"},
				{Title: "RRR", Director: "Rajamouli"},
			},
		}
		tmpl.Execute(w,films)
	}
	h2 := func (w http.ResponseWriter, r *http.Request){
		log.Println("HTMX request received")
		log.Println(r.Header.Get("HX-Request"))
	} 
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film",h2)
	log.Fatal(http.ListenAndServe(":8000",nil))
}