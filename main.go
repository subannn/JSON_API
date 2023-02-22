package main

import (
	//"bytes"
	// "database/sql"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"
	_ "github.com/lib/pq"

	//"golang.org/x/crypto/bcrypt"

	hndl "./handlers"
	db "./DB"

)

func main() {

	db.OpenDB()
	defer db.DB.Close()

	http.HandleFunc("/jsonPost", hndl.JsonPost) 

	http.HandleFunc("/tableGet", hndl.TableGet)

	http.HandleFunc("/jsonGetByID", hndl.JsonGet)

	http.HandleFunc("/jsonPut", hndl.JsonPut)

	http.HandleFunc("/jsonDelate", hndl.JsonDelate) 

	http.ListenAndServe(":8080", nil)
}	