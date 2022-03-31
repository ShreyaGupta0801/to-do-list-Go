package middleware

import (
	"encoding/json"
	"fmt"
	"golang-react-to-do/server/database"
	"golang-react-to-do/server/models"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(r.Method)
	if r.Method == "OPTIONS" {
		return
	}
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("create hit with user", user)
	if user.Username == "" {
		return
	}
	var addedUser = database.InsertUser(user)
	json.NewEncoder(w).Encode(addedUser)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("login hit with user", user)
	var loggedUser = database.CheckUserLogin(user)
	json.NewEncoder(w).Encode(loggedUser)
}
