package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type language struct {
	Title string
	Lang1 string
	Lang2 string
	Lang3 string
}

func homepage(w http.ResponseWriter, r *http.Request) {
	// html := "Welcome to Golang!"
	// fmt.Fprint(w, html)

	Language := language{
		Title: "Learning of Programming languages",
		Lang1: "Golang",
		Lang2: "Python",
		Lang3: "Java",
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, Language)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", homepage)
	log.Println("Starting Webserver on port 8080")
	http.ListenAndServe(":8080", nil)
}
