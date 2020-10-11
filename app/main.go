package main

import (
	"log"
	"net/http"

	"github.com/deozza/go-blog-native/utils"
	"github.com/deozza/go-blog-native/user"
)

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&User{})
}

func main() {
	db := utils.InitDb()
	Migrate(db)
	defer db.Close()

	r := utils.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

