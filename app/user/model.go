package user

import (
	"time"
	"gorm.io/gorm"
)

type User struct {

	Uuid      string    `json:"uuid" gorm:"primary_key"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsActive  bool      `json:"isActive" gorm:"default:0"`
	Roles     string    `json:"roles" gorm:"default:"{"ROLE_USER"}"`
}


func Save(u *User) {

}

func FindAll() {
}

func FindByUuid(uuid string){

}

func Delete(uuid string) {
}