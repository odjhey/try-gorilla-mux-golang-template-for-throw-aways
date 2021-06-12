package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the homepage")
	fmt.Println("endpoitnn hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		{Title: "t1", Desc: "d", Content: "content123213"},
		{Title: "t12", Desc: "d1", Content: "content123213asdf"},
	}
	handleRequest()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Articles hit")
	json.NewEncoder(w).Encode(Articles)
}

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
