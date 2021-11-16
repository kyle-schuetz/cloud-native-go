package main

import (
	"log"
	"net/http"

	"github.com/cloud-native-go/key_value_store/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/{key}", handler.KeyValuePutHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
