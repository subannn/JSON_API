package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

	_"github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"

	db "dependencies/db"
)

type UserUpd struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Mail     string `json:"mail"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func TableGet(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT * FROM users")
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
			CheckError(err)
		}
		buf.WriteString(fmt.Sprintf("%d: %s: %s: %s: %s: %s\n", id, name, surname, mail, pnumber, pass))
	}
	if err := rows.Err(); err != nil {
		CheckError(err)
	}

	fmt.Fprint(w, buf.String())
}


func JsonGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    id := params["id"]

	psqlInfo := fmt.Sprintf("SELECT * FROM users WHERE id=%s", id)

	rows, err := db.DB.Query(psqlInfo)
	CheckError(err)

	var u User

	for rows.Next() {
		var name string
    	var surname string
		var mail string
		var phone string
		var password string
		if err := rows.Scan(&id, &name, &surname, &mail, &phone, &password); err != nil {
			CheckError(err)
		}
		u.Name = name
		u.Surname = surname
		u.Mail = mail
		u.Phone = phone
		u.Password = password
	
	}
	
	jsonU,err := json.Marshal(u)
	CheckError(err)
	
	fmt.Fprint(w, string(jsonU))	
}

func JsonPost(w http.ResponseWriter, r *http.Request) {

	var u User

	json.NewDecoder(r.Body).Decode(&u)

	CryptedHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	CheckError(err)

	_, err = db.DB.Exec("INSERT INTO users(name, surname, mail, phone, password) VALUES ($1, $2, $3, $4, $5)", u.Name, u.Surname, u.Mail, u.Phone, CryptedHash)
	CheckError(err)
}
func JsonPut(w http.ResponseWriter, r *http.Request) {

	var u UserUpd

	json.NewDecoder(r.Body).Decode(&u)

	CryptedHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	CheckError(err)

	_, err = db.DB.Exec("UPDATE users SET name=$2 , surname=$3, mail=$4 , phone=$5 , password=$6 WHERE id=$1", u.Id, u.Name, u.Surname, u.Mail, u.Phone, CryptedHash)
	if err != nil {
		panic(err)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    id := params["id"]

	_, err := db.DB.Exec("DELETE FROM users WHERE id = $1", id)
	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}