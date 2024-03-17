package web

type PenyakitCreated struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ObatCreated struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"string"`
	Price        int32  `json:"price"`
	QtyAvailable uint8  `json:"qty_available"`
	Composition  string `json:"composition"`
}

type ObatUpdated struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"string"`
	Price        int32  `json:"price"`
	QtyAvailable uint8  `json:"qty_available"`
}