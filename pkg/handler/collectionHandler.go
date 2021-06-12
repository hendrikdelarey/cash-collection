package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type OwedMoneyDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phoneNumber"`
	Email             string `json:"email"`
	AmountOwedInCents int    `json:"amountOwedInCents"`
	AmountPaidInCents int    `json:"amountPaidInCents"`
	ReferenceCode     string `json:"referenceCode"`
	CreatedDate       string `json:"createdDate"`
}

type CollectionDTO struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	OwedMoney   []OwedMoneyDTO `json:"owedMoney"`
	CreatedDate string         `json:"dateCreated"`
}

var collections []CollectionDTO

func DeleteCollection(w http.ResponseWriter, r *http.Request) {
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

func UpdateCollection(w http.ResponseWriter, r *http.Request) {
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
	http.Error(w, "The requested item to update was not found.", http.StatusNotFound)
}

func CreateCollection(w http.ResponseWriter, r *http.Request) {
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

func GetCollection(w http.ResponseWriter, r *http.Request) {
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

func GetCollections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(collections) == 0 {
		json.Marshal(make([]*CollectionDTO, 0))
		return
	}
	json.NewEncoder(w).Encode(collections)
}
