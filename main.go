package main

import (
	"fmt"
	"log"
	"net/http"

	"myapis/main/handlers"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the homepage")
	fmt.Println("endpoitnn hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	jsonSubRoutes := myRouter.PathPrefix("/api").Subrouter()
	myRouter.HandleFunc("/", homePage)

	jsonSubRoutes.HandleFunc("/articles", handlers.ReturnAllArticles)
	jsonSubRoutes.HandleFunc("/article", handlers.CreateNewArticle).Methods("POST")
	jsonSubRoutes.HandleFunc("/article/{id}", handlers.ReturnSingleArticle)
	jsonSubRoutes.Use(loggingMiddleware)

	myRouter.HandleFunc("/hello", returnHello).Methods("POST")
	myRouter.HandleFunc("/hello", returnHello)
	myRouter.HandleFunc("/sss", returnSSS)

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

	handlers.Articles = []handlers.Article{
		{Id: "1", Title: "t1", Desc: "d", Content: "content123213"},
		{Id: "2", Title: "t12", Desc: "d1", Content: "content123213asdf"},
	}
	handleRequest()
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
