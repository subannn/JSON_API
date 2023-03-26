package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	db "suba/dependencies/db"
	rgst "suba/dependencies/requests"
)


func main() {

	db.OpenDB()
	defer db.DB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/jsonPost", rgst.JsonPost).Methods("POST")
	router.HandleFunc("/tableGet", rgst.TableGet).Methods("GET")
	router.HandleFunc("/jsonGet/{id}", rgst.JsonGet).Methods("GET")
	router.HandleFunc("/jsonPut", rgst.JsonPut).Methods("PUT")
	router.HandleFunc("/Delete/{id}", rgst.Delete).Methods("DELETE")

	server := &http.Server {
		Handler:      router,
		Addr: 		  "127.0.0.1:8000",
		ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	
	log.Fatal(server.ListenAndServe())

}
