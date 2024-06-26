package web

type Penyakit struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Obat        []Obat `json:"obat"`
}

type Obat struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	QtyAvailable int    `json:"qty_available"`
	QtyReserved  int    `json:"qty_reserved"`
	QtyOn        int    `json:"qty_on"`
	Composition  string `json:"composition"`
}
