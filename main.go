package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"myapis/main/docs"
	_ "myapis/main/docs"
	"myapis/main/handlers"
	"myapis/main/store"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

var (
	buildTime  string
	commitHash string
	yow        string
)

// @title ze golang api
// @version 1.0
// @description sampol
// @description build
// @contact.email odjhey@gmail.com
// @BasePath /
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	fmt.Printf("Build Time: %s ", buildTime)
	fmt.Printf("Commit Hash: %s %s\n", commitHash, os.Getenv("DOKKU_GIT_REV"))
	docs.SwaggerInfo.Description = fmt.Sprintf(`Sampol
	buildTime: %s commit: %s %s`, buildTime, commitHash, os.Getenv("DOKKU_GIT_REV"))

	handlers.Articles = []handlers.Article{
		{Id: "1", Title: "t1", Desc: "d", Content: "content123213"},
		{Id: "2", Title: "t12", Desc: "d1", Content: "content123213asdf"},
	}

	handlers.XMLArticles = []handlers.XMLArticle{
		{Id: "1", Title: "t1", Desc: "d", Content: "content123213"},
		{Id: "2", Title: "t12", Desc: "d1", Content: "content123213asdf"},
	}
	handleRequest()
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(loggingMiddleware)

	// json routes
	jsonSubRoutes := myRouter.PathPrefix("/api").Subrouter()
	jsonSubRoutes.Use(jsonMiddleware)

	jsonSubRoutes.HandleFunc("/articles", handlers.ReturnAllArticles)
	jsonSubRoutes.HandleFunc("/article", handlers.CreateNewArticle).Methods("POST")
	jsonSubRoutes.HandleFunc("/article/{id}", handlers.ReturnSingleArticle)

	jsonSubRoutes.HandleFunc("/product", store.CreateProductHandler).Methods("POST")
	jsonSubRoutes.HandleFunc("/product/{id}", store.GetSingleProduct)
	jsonSubRoutes.HandleFunc("/products", store.GetAllProductsHandler)

	// xml routes
	xmlSubRoutes := myRouter.PathPrefix("/xml").Subrouter()
	xmlSubRoutes.Use(xmlMiddleware)

	xmlSubRoutes.HandleFunc("/articles", handlers.ReturnAllXMLArticles)
	xmlSubRoutes.HandleFunc("/article", handlers.CreateNewXMLArticle).Methods("POST")
	xmlSubRoutes.HandleFunc("/article/{id}", handlers.ReturnSingleXMLArticle)

	myRouter.HandleFunc("/hello", handlers.ReturnHello).Methods("POST")
	myRouter.HandleFunc("/hello", handlers.ReturnHello)
	myRouter.HandleFunc("/sss", handlers.ReturnSSS)

	myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	myRouter.HandleFunc("/", handlers.HomePage)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func xmlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprint(w, xml.Header)
		next.ServeHTTP(w, r)
	})
}
