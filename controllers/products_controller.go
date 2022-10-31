package controllers

import (
	"encoding/json"
	"net/http"
	"store-manager/services"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	response, err := services.GetProducts()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("Error in JSON encoding"))
		return
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Error converting ID to int"))
	}

	response, err := services.GetProductById(ID)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("Error in JSON encoding"))
		return
	}
}