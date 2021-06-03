package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OwedMoneyDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phoneNumber"`
	Email             string `json:"email"`
	AmountOwedInCents int    `json:"amountOwedInCents"`
	AmountPaidInCents int    `json:"amountPaidInCents"`
	ReferenceCode     string `json:"referenceCode"`
	CreatedDate       string `json: "createdDate"`
}
type CollectionDTO struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	OwedMoney   []OwedMoneyDTO `json:"owedMoney"`
	CreatedDate string         `json: "dateCreated"`
}

var collections []CollectionDTO

func deleteCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range collections {
		if item.ID == params["id"] {
			collections = append(collections[:index], collections[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(collections)
}

func updateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range collections {
		if item.ID == params["id"] {
			collections = append(collections[:index], collections[index+1:]...)
			var c CollectionDTO
			_ = json.NewDecoder(r.Body).Decode(&c)
			c.ID = params["id"]
			collections = append(collections, c)
			json.NewEncoder(w).Encode(c)
			return
		}
	}

}

func createCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var c CollectionDTO
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.ID = strconv.Itoa(rand.Intn(1000000))
	collections = append(collections, c)
	json.NewEncoder(w).Encode(c)
}

func getCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range collections {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func getCollections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(collections) == 0 {
		json.Marshal(make([]*CollectionDTO, 0))
		return
	}
	json.NewEncoder(w).Encode(collections)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/collections", getCollections).Methods("GET")
	router.HandleFunc("/collections", createCollection).Methods("POST")
	router.HandleFunc("/collections/{id}", getCollection).Methods("GET")
	router.HandleFunc("/collections/{id}", updateCollection).Methods("PUT")
	router.HandleFunc("/collections/{id}", deleteCollection).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
