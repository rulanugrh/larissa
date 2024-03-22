package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FName    string    `json:"fname" form:"fname"`
	LName    string    `json:"lname" form:"lname"`
	Email    string    `json:"email" form:"email"`
	TTL      time.Time `json:"ttl" form:"ttl"`
	Password string    `json:"password" form:"password"`
	Address  string    `json:"address" form:"address"`
	Age      int       `json:"age" form:"age"`
	RoleID   uint      `json:"role_id" form:"role_id"`
	Role     Role      `json:"role" form:"role" gorm:"foreignKey:RoleID;references:ID"`
}

type UserRegister struct {
	FName    string    `json:"fname" form:"fname" validate:"required"`
	LName    string    `json:"lname" form:"lname" validate:"required"`
	Email    string    `json:"email" form:"email" validate:"required,email"`
	TTL      time.Time `json:"ttl" form:"ttl" validate:"required"`
	Age      int       `json:"age" form:"age" validate:"required"`
	Password string    `json:"password" form:"password" validate:"required,min=8"`
	Address  string    `json:"address" form:"address" validate:"required"`
	RoleID   uint      `json:"role_id" form:"role_id" validate:"required"`
}

type UserLogin struct {
	FName    string `json:"fname" form:"fname"`
	LName    string `json:"lname" form:"lname"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Role struct {
	gorm.Model
	Name         string `json:"name"`
	Descriptions string `json:"desc"`
}
