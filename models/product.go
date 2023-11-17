package models

import "time"

type User struct {
	Username  string    `json:"username" bson:"_id"`
	Password  string    `json:"password" bson:"password"`
	Roles     []string  `json:"roles" bson:"roles"`
	LastLogin time.Time `json:"lastLogin" bson:"lastLogin"`
	IsActive  bool      `json:"isActive" bson:"isActive"`
}