package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "deozza/golang-blog-native/endpoints"
)

func buildRouter() http.Handler {
	myRouter := mux.NewRouter().StrictSlash(true)
    
    myRouter.HandleFunc("/test", endpoints.ReturnTest).Methods("GET")

    myRouter.HandleFunc("/users", endpoints.GetUserList).Methods("GET")
    myRouter.HandleFunc("/users", endpoints.PostUser).Methods("POST")
    myRouter.HandleFunc("/users/{id}", endpoints.GetUser).Methods("GET")
    myRouter.HandleFunc("/users/{id}", endpoints.PatchUser).Methods("PATCH")
    myRouter.HandleFunc("/users/{id}", endpoints.DeleteUser).Methods("DELETE")


    return myRouter
}