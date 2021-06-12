package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `json:"Code"`
	Price uint   `json:"Price"`
}

func Connect() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(".data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	// migrate schema
	db.AutoMigrate(&Product{})

	/*
		// read
		var product Product
		db.First(&product, 1)
		result := db.First(&product, "code = ?", "123")
		fmt.Println("hmm")
		fmt.Println(result)

		// update
		db.Model(&product).Update("Price", 499)
	*/

	return

}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	var products []Product
	result := db.Find(&products)
	fmt.Println(result)
	// fmt.Fprint(w, "Keut!")
	json.NewEncoder(w).Encode(products)

}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	body, _ := ioutil.ReadAll(r.Body)
	var product Product
	err := json.Unmarshal(body, &product)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Print(string(body))

	db.Create(&product)
	/*
		fmt.Printf("%+v", product)
	*/

	log, _ := json.MarshalIndent(product, "", " ")
	fmt.Print(string(log))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)

}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["id"]

	db := Connect()
	var product Product
	tx := db.First(&product, "Code = ?", code)
	if tx.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log, _ := json.MarshalIndent(product, "", " ")
	fmt.Print(string(log))

	json.NewEncoder(w).Encode(product)

}
