package controllers

import (
	"encoding/json"
	"net/http"
	"store-manager/services"
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