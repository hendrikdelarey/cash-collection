package router

import (
	"github.com/gorilla/mux"
	"github.com/hendrikdelarey/cash-collection/pkg/handler"
)

func New() {
	router := mux.NewRouter()

	router.HandleFunc("/collections", handler.GetColletion).Methods("GET")
	router.HandleFunc("/collections", handler.CreateCollection).Methods("POST")
	router.HandleFunc("/collections/{id}", handler.GetCollection).Methods("GET")
	router.HandleFunc("/collections/{id}", handler.UpdateCollection).Methods("PUT")
	router.HandleFunc("/collections/{id}", handler.DeleteCollection).Methods("DELETE")

	router.HandleFunc("/transactions", handler.GetRecentInvestecTransactions).Methods("GET")

	router.HandleFunc("/login", handler.LoginUser).Methods("POST")
	router.HandleFunc("/register", handler.RegisterUser).Methods("POST")
}
