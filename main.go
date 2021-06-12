package main

import (
	"fmt"
	"log"
	"net/http"

	"myapis/main/handlers"
	"myapis/main/store"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	handlers.Articles = []handlers.Article{
		{Id: "1", Title: "t1", Desc: "d", Content: "content123213"},
		{Id: "2", Title: "t12", Desc: "d1", Content: "content123213asdf"},
	}
	handleRequest()
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	jsonSubRoutes := myRouter.PathPrefix("/api").Subrouter()
	myRouter.HandleFunc("/", handlers.HomePage)

	jsonSubRoutes.HandleFunc("/articles", handlers.ReturnAllArticles)
	jsonSubRoutes.HandleFunc("/article", handlers.CreateNewArticle).Methods("POST")
	jsonSubRoutes.HandleFunc("/article/{id}", handlers.ReturnSingleArticle)

	jsonSubRoutes.HandleFunc("/product", store.CreateProductHandler).Methods("POST")
	jsonSubRoutes.HandleFunc("/product/{id}", store.GetSingleProduct)
	jsonSubRoutes.HandleFunc("/products", store.GetAllProductsHandler)
	jsonSubRoutes.Use(loggingMiddleware)

	myRouter.HandleFunc("/hello", handlers.ReturnHello).Methods("POST")
	myRouter.HandleFunc("/hello", handlers.ReturnHello)
	myRouter.HandleFunc("/sss", handlers.ReturnSSS)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
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
