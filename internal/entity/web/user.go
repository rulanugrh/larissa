package web

import "time"

type User struct {
	ID      uint      `json:"id"`
	Email   string    `json:"email"`
	FName   string    `json:"fname"`
	LName   string    `json:"lname"`
	Age     int       `json:"age"`
	Address string    `json:"address"`
	TTL     time.Time `json:"ttl"`
}

type Kunjungan struct {
	FName    string     `json:"fname"`
	LName    string     `json:"lname"`
	Age      int        `json:"age"`
	Address  string     `json:"address"`
	Penyakit []Penyakit `json:"penyakit"`
}
