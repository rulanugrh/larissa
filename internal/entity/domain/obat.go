package domain

import "gorm.io/gorm"

type Obat struct {
	gorm.Model
	Name         string     `json:"name" form:"name"`
	Description  string     `json:"description" form:"description"`
	Price        int32      `json:"price" form:"price"`
	QtyOn        uint8      `json:"qty_on" form:"qty_on"`
	QtyAvailable uint8      `json:"qty_available" form:"qty_available"`
	QtyReserved  uint8      `json:"qty_reversed" form:"qty_reversed"`
	Composition  string     `json:"composition"`
	Penyakit     []Penyakit `json:"penyakit" form:"penyakit" gorm:"many2many:list_penyakit"`
}
