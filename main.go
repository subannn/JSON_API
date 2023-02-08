package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "go_crud"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func second_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/suban_page.html")
	CheckError(err)
	
	tmpl.Execute(w, nil)
}

func delate(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	DelateInDB(id)
}
func DelateInDB(id string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	CheckError(nil)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	CheckError(err)
}

func save(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("fname")
	surname := r.FormValue("lname")
	mail := r.FormValue("mail")
	phone := r.FormValue("phone")
	pass := r.FormValue("pass")
	SaveInDB(name, surname, mail, phone, pass)
}

func SaveInDB(name, surname, mail, phone, pass string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	CheckError(nil)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	_, err = db.Exec("INSERT INTO users(fname, lname, mail, pnumber, pass) VALUES ($1, $2, $3, $4, $5)", name, surname, mail, phone, pass)
	CheckError(err)
}

func showtable(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	CheckError(nil)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	rows, err := db.Query("SELECT * FROM users")
	CheckError(err)

	defer rows.Close()

	
	var buf bytes.Buffer
	for rows.Next() {
		var id int
    	var name string
    	var surname string
		var mail string
		var pnumber string
		var pass string
    	if err := rows.Scan(&id, &name, &surname, &mail, &pnumber, &pass); err != nil {
       		log.Fatal(err)
    	}
    	buf.WriteString(fmt.Sprintf("%d: %s: %s: %s: %s: %s\n", id, name, surname, mail, pnumber, pass))
	}
	if err := rows.Err(); err != nil {
    	log.Fatal(err)
	}
	fmt.Fprint(w, buf.String())
}

func edit(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("fname")
	surname := r.FormValue("lname")
	mail := r.FormValue("mail")
	phone := r.FormValue("phone")
	pass := r.FormValue("pass")
	EditDB(id, name, surname, mail, phone, pass)
}

func EditDB(id, name, surname, mail, phone, pass string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	CheckError(nil)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	if(name != "") {
		_, err = db.Exec("UPDATE users SET fname=$2 WHERE id=$1", id, name)
		CheckError(err)
	}

	if(surname != "") {
		_, err = db.Exec("UPDATE users SET lname=$2 WHERE id=$1", id, surname)
		CheckError(err)
	}

	if(mail != "") {
		_, err = db.Exec("UPDATE users SET mail=$2 WHERE id=$1", id, mail)
		CheckError(err)
	}

	if(phone != "") {
		_, err = db.Exec("UPDATE users SET phone=$2 WHERE id=$1", id, phone)
		CheckError(err)
	}

	if(pass != "") {
		_, err = db.Exec("UPDATE users SET pass=$2 WHERE id=$1", id, pass)
		CheckError(err)
	}
}






func handleRequest() {
	http.HandleFunc("/", home_page)

	http.HandleFunc("/second_page/", second_page)
	http.HandleFunc("/second_page", second_page)

	http.HandleFunc("/save", save)

	http.HandleFunc("/delate", delate)

	http.HandleFunc("/showtable", showtable)

	http.HandleFunc("/edit", edit)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()

}
func CheckError(err error) {
	if(err != nil) {
		panic(err)
	}
}