package models

import (
	"time"
	"encoding/json"
)

type User struct {
	Username string `json:Username`
	Email string `json:email`
	Password string 
	CreatedAt time.Time `json:createdAt`
	UpdatedAt time.Time `json:updatedAt`
	Roles []string `json:roles`
	IsActive bool `json:isActive`
}

