package main

import (
	"net/http"
	"log"
	"gorm.io/gorm"
	"github.com/gorilla/mux"

	"github.com/deozza/golang-blog-native/commons"
	"github.com/deozza/golang-blog-native/user"
)

type App struct {
	Db *gorm.DB
	Router *mux.Router
}

func Migrate(app App){
	app.Db.Debug().AutoMigrate(&user.User{})
}


func main() {

	app := App{}

	app.Db = commons.InitDb()
	Migrate(app)
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/api/users", user.GetUsers).Methods("GET")
	app.Router.HandleFunc("/api/users", user.PostUser).Methods("POST")


	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

