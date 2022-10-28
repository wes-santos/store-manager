package routes

import (
	"net/http"
	"store-manager/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", controllers.GetProducts).Methods(http.MethodGet)

	return router
}