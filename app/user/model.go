package user

import (
	"time"
	guuid "github.com/google/uuid"
)

type User struct {
	ID        uint       `json:"-"         gorm:"primary_key;column:id"`
	Uuid      guuid.UUID `json:"uuid"      gorm:"column:uuid;unique;not null"`
	Username  string     `json:"username"  gorm:"column:username;size:255;unique;not null"`
	Email     string     `json:"email"     gorm:"column:email;size:255;unique;not null"`
	Password  string     `json:"password"  gorm:"column:password;not null"`
	CreatedAt time.Time  `json:"CreatedAt" gorm:"column:createdAt;autoCreateTime;not null"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime;not null"`
	IsActive  bool       `json:"isActive"  gorm:"column:isActive;default:0"`
	Roles     string     `json:"roles"     gorm:"column:roles`
}

func NewUser() (user User){
	user = User{Uuid: guuid.New(),Roles:`{"ROLE_USER"}`}
	return
}