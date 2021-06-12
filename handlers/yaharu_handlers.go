package handlers

import (
	"fmt"
	"net/http"
)

func GetYaharu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yaharu!")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the homepage")
	fmt.Println("endpoitnn hit")
}

func ReturnSSS(w http.ResponseWriter, r *http.Request) {
	const asdf = `asdfasdf`
	fmt.Fprint(w, asdf)
}

func ReturnHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"yo": "brow"}`)
}
