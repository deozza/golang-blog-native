package user

import (
	"net/http"
	"encoding/json"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}