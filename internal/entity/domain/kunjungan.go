package domain

import "gorm.io/gorm"

type Kunjungan struct {
	gorm.Model
	UserID     uint       `json:"user_id" form:"user_id" validate:"required"`
	PenyakitID []uint     `json:"penyakit_id" form:"penyakit_id" validate:"required"`
	User       User       `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Penyakit   []Penyakit `json:"penyakit" gorm:"foreignKey:PenyakitID;references:ID"`
}
