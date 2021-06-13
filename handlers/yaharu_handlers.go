package handlers

import (
	"fmt"
	"net/http"
)

func GetYaharu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yaharu!")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "welcome to the homepage <br>")
	fmt.Fprint(w, `go check the <a href="/swagger/index.html">swagger UI</a>`)
	fmt.Println("endpoint hit")
}

func ReturnSSS(w http.ResponseWriter, r *http.Request) {
	const asdf = `asdfasdf`
	fmt.Fprint(w, asdf)
}

func ReturnHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"yo": "brow"}`)
}
