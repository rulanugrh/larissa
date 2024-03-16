package domain

import "gorm.io/gorm"

type Penyakit struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"desc" form:"desc"`
	ObatID      []uint `json:"obat_id" form:"obat_id"`
	Obat        []Obat `json:"obat" form:"obat" gorm:"foreignKey:ObatID;reference:ID;many2many:resep"`
}
