package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
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
		time.Sleep(2*time.Second)
		// log.Println("HTMX request received")
		// log.Println(r.Header.Get("HX-Request"))
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		// fmt.Println(title, director)

		htmlStr := fmt.Sprintf("<li class='bg-blue-500 text-white p-2 mb-2 rounded-md'> %s - %s </li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, tmpl)
	} 
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film",h2)
	log.Fatal(http.ListenAndServe(":8000",nil))
}