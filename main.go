package main

import (
	"fmt"
	"log"
	"net/http"
	"store-manager/routes"
)

func main() {
	router := routes.Routes()

	fmt.Println("Listening: 3500")

	log.Fatal(http.ListenAndServe(":3500", router))
}
