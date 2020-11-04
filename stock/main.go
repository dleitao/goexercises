package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var stock map[int]product

type product struct {
	ID    int     `bson:"id,omitempty" json:"id,omitempty"`
	Name  string  `bson:"name,omitempty" json:"name,omitempty"`
	Count int     `bson:"count,omitempty" json:"count,omitempty"`
	Price float32 `bson:"price,omitempty" json:"price,omitempty"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("API is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var name string
	id := len(stock) + 1
	p := product{id, name, 0, 0.0}

	json.NewDecoder(r.Body).Decode(&p)
	stock[id] = p

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list := []product{}
	for _, elem := range stock {
		list = append(list, elem)
	}

	response, err := json.Marshal(list)
	if err != nil {
		log.Println(err)
	}

	io.WriteString(w, string(response))
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	response, err := json.Marshal(stock[id])
	if err != nil {
		log.Println(err)
	}

	io.WriteString(w, string(response))
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, ok := stock[id]
	if ok {
		delete(stock, id)
		io.WriteString(w, `{"deleted": true}`)
	} else {
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p, ok := stock[id]

	if ok {
		json.NewDecoder(r.Body).Decode(&p)
		stock[id] = p
		res, _ := json.Marshal(p)
		io.WriteString(w, string(res))
	} else {
		io.WriteString(w, `{"update": false, "error": "Record Not Found"}`)
	}
}

// func updateProductField(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])
// 	p, ok := stock[id]

// 	if ok {
// 		p1 := product{}
// 		json.NewDecoder(r.Body).Decode(&p1)
// 		log.Println(p1)
// 		res, _ := json.Marshal(p)
// 		io.WriteString(w, string(res))
// 	} else {
// 		io.WriteString(w, `{"update": false, "error": "Record Not Found"}`)
// 	}
// }

func init() {
	stock = make(map[int]product, 100)
	stock[1] = product{1, "firstProduct", 87, 1.25}
}

func main() {
	log.Println("Starting API server")
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/all", getAllProducts).Methods("GET")
	router.HandleFunc("/product", createProduct).Methods("POST")
	router.HandleFunc("/product/{id}", getProduct).Methods("GET")
	router.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
	// router.HandleFunc("/product/{id}", updateProductField).Methods("PATCH")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
	}).Handler(router)
	http.ListenAndServe(":8000", handler)

}

// curl -H "Content-Type: application/json" --data '{"name": "2ndProduct", "price": 122.99}' http://localhost:8000/product
// curl -X PATCH -H "Content-Type: application/json" --data '{"name": "1stProduct"}' http://localhost:8000/product/1
// curl -X DELETE http://localhost:8000/product/1
// curl -X GET http://localhost:8000/product/all
