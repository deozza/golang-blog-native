package user

import (
	"net/http"
	"encoding/json"

	"github.com/deozza/golang-blog-native/commons"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := NewUser()
 
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer r.Body.Close()

	if err := commons.GetDb().Save(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

/*

	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)
	*/
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}