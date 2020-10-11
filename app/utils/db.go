package utils

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var db *gorm.DB

func InitDb() *gorm.DB{
	d, err := gorm.Open(mysql.Open("go:go@tcp(mysql_go_blog_native:3306)/go-blog-native?charset=utf8&parseTime=True"), &gorm.Config{})

 	if err!=nil{
 		panic(err.Error())
 	} 

 	db = d

 	return db
}

func GetDb()  *gorm.DB{
	return db
}