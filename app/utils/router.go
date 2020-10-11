package utils

import (
	"github.com/gorilla/mux"

	"github.com/deozza/golang-blog-native/user"
)

func InitRouter() (r *mux.Router){
	r = mux.NewRouter()
	r.HandleFunc("/api/users"     , user.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", user.GetUser).Methods("GET")
	r.HandleFunc("/api/users"     , user.PostUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", user.PatchUser).Methods("PATCH")
	r.HandleFunc("/api/users/{id}", user.DeleteUser).Methods("DELETE")

	return
}