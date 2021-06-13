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
	Code  string `json:"Code" example:"X909" validate:"required"` // Ze product code
	Price uint   `json:"Price" example:"99"`                      // No decimals for now
}

type ResponseMessage struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
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

// GetAllProducts godoc
// @Summary Get details of all products
// @Description Get details of all products
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {object} []Product
// @Router /api/products [get]
func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	var products []Product
	result := db.Find(&products)
	fmt.Println(result)
	// fmt.Fprint(w, "Keut!")
	json.NewEncoder(w).Encode(products)

}

// CreateProduct godoc
// @Summary Create product
// @Description Create product
// @Tags product
// @Accept  json
// @Produce  json
// @Param body body Product true "input payload"
// @Success 201 {object} Product
// @Router /api/product [post]
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	body, _ := ioutil.ReadAll(r.Body)
	var product Product
	err := json.Unmarshal(body, &product)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ResponseMessage{Message: "Bad payload.", Code: ""})
		return
	}

	var existingProduct Product
	res := db.First(&existingProduct, "Code = ?", product.Code)
	fmt.Println(res.RowsAffected, product.Code)
	if res.Error == nil {
		fmt.Println("Already exist.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ResponseMessage{Message: "Resource already exist.", Code: ""})
		return
	}

	db.Create(&product)
	/*
		fmt.Printf("%+v", product)
	*/

	log, _ := json.MarshalIndent(product, "", " ")
	fmt.Print(string(log))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)

}

// GetProduct godoc
// @Summary Read 1 product
// @Description Get 1 product
// @Tags product
// @Accept  json
// @Produce  json
// @Param code path string true "product code"
// @Success 200 {object} Product
// @Failure 404 {object} ResponseMessage
// @Router /api/product/{code} [post]
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
