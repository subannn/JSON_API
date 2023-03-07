package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	db "suba/example/db"
	rgst "suba/example/requests"
)

func main() {

	db.OpenDB()
	defer db.DB.Close()

	r := mux.NewRouter()


	r.HandleFunc("/jsonPost", rgst.JsonPost).Methods("POST")
	r.HandleFunc("/tableGet", rgst.TableGet).Methods("GET")
	r.HandleFunc("/jsonGet/{id}", rgst.JsonGet).Methods("GET")
	r.HandleFunc("/jsonPut", rgst.JsonPut).Methods("PUT")
	r.HandleFunc("/Delete/{id}", rgst.Delete).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", r))

}
