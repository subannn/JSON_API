package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"os"
    "os/signal"
	"flag"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	db "dependencies/db"
	rgst "dependencies/requests"
)


func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()	

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
		
	}
	
	log.Fatal(server.ListenAndServe())

	c := make(chan os.Signal, 1)
    
    signal.Notify(c, os.Interrupt)

    <-c

    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    
    server.Shutdown(ctx)
    
    log.Println("shutting down")
    os.Exit(0)

}
