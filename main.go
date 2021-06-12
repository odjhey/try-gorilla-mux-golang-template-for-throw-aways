package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the homepage")
	fmt.Println("endpoitnn hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/hello", returnHello).Methods("POST")
	myRouter.HandleFunc("/hello", returnHello)
	myRouter.HandleFunc("/sss", returnSSS)

	myRouter.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSSS(w http.ResponseWriter, r *http.Request) {
	const asdf = `asdfasdf`
	fmt.Fprint(w, asdf)
}

func returnHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"yo": "brow"}`)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		{Id: "1", Title: "t1", Desc: "d", Content: "content123213"},
		{Id: "2", Title: "t12", Desc: "d1", Content: "content123213asdf"},
	}
	handleRequest()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Articles hit")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println(vars)
	fmt.Fprint(w, "Key: "+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
