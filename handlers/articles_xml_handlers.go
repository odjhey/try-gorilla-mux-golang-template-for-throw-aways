package handlers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type XMLArticle struct {
	XMLName  xml.Name `xml:"Article"`
	SomeAttr int      `xml:"data-id,attr"`
	Id       string   `xml:"id"`
	Title    string   `xml:"title"`
	Desc     string   `xml:"desc"`
	Content  string   `xml:"content"`
}

type TXMLArticles struct {
	XMLName  xml.Name `xml:"Articles"`
	Articles []XMLArticle
}

var XMLArticles []XMLArticle

func ReturnAllXMLArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("XML Articles hit")
	xml.NewEncoder(w).Encode(&TXMLArticles{Articles: XMLArticles})
}

func ReturnSingleXMLArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println(vars)
	fmt.Fprint(w, "Key: "+key)
	for _, article := range XMLArticles {
		if article.Id == key {
			xml.NewEncoder(w).Encode(article)
		}
	}
}

func CreateNewXMLArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article XMLArticle
	xml.Unmarshal(reqBody, &article)

	fmt.Printf("%+v", article)

	XMLArticles = append(XMLArticles, article)

	xml.NewEncoder(w).Encode(article)
}
