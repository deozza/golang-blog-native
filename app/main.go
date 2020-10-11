package main

import (
	"net/http"
	"log"
	"gorm.io/gorm"

	"github.com/deozza/golang-blog-native/commons"
	"github.com/deozza/golang-blog-native/user"
)

func Migrate(db *gorm.DB){
	db.Debug().AutoMigrate(&user.User{})
}


func main() {
	db := commons.InitDb()
	Migrate(db)
	r := commons.BuildRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}

