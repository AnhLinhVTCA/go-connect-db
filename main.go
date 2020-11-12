package main

import (
	"connect-to-mysql/dal"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/item", dal.GetAllItems).Methods(http.MethodGet)
	r.HandleFunc("/api/item/{id}", dal.GetItemByID).Methods(http.MethodGet)
	r.HandleFunc("/api/item", dal.CreateItem).Methods(http.MethodPost)
	r.HandleFunc("/api/item/{id}", dal.UpdateItem).Methods(http.MethodPut)
	r.HandleFunc("/api/item/{id}", dal.DeleteItem).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}
