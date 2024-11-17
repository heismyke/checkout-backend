package main

import (
	"log"
	"net/http"

	"github.com/heismyke/checkout_backend/handler"
)

func main(){

	http.HandleFunc("/products", handler.HandleCreatePaymentIntent)

	var err error = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}