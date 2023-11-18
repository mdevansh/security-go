package models

import "time"

type User struct {
	Username  string    `json:"username" bson:"_id" binding:"required"`
	Password  string    `json:"password" bson:"password" binding:"required"`
	Roles     []string  `json:"roles" bson:"roles" binding:"required"`
	LastLogin time.Time `json:"lastLogin" bson:"lastLogin"`
	IsActive  bool      `json:"isActive" bson:"isActive"`
}
