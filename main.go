package main

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	body := "Hello World"

	page := &Page{Title:title, Body:body}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}
