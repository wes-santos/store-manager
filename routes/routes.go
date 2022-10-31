package routes

import (
	"net/http"
	"store-manager/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	// GET ALL PRODUCTS
	router.HandleFunc("/products", controllers.GetProducts).Methods(http.MethodGet)

	// GET PRODUCT BY ID
	router.HandleFunc("/products/{id}", controllers.GetProductById).Methods(http.MethodGet)

	return router
}