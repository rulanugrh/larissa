package domain

import "gorm.io/gorm"

type Kunjungan struct {
	gorm.Model
	UserID     uint              `json:"user_id" form:"user_id" validate:"required"`
	Doctor     string            `json:"doctor" from:"doctor" validate:"required"`
	Type       uint              `json:"type" form:"type"`
	PenyakitID []uint            `json:"penyakit_id" form:"penyakit_id" validate:"required"`
	User       User              `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Penyakit   []Penyakit        `json:"penyakit" gorm:"foreignKey:PenyakitID;references:ID"`
	Category   CategoryKunjungan `json:"category" form:"category" gorm:"foreignKey:Type;references:ID"`
}

type CategoryKunjungan struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}
