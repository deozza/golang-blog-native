package commons

import (
	"github.com/gorilla/mux"

	"github.com/deozza/golang-blog-native/user"
)

func BuildRouter() (r *mux.Router) {
	r = mux.NewRouter()

	r.HandleFunc("/api/users", user.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", user.PostUser).Methods("POST")
	return
}
