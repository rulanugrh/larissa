package config

import (
	"fmt"
	"log"

	"github.com/rulanugrh/larissa/internal/entity/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
	conf *App
}

func InitializePostgres(conf *App) *Postgres {
	return &Postgres{conf: conf}
}

func (p *Postgres) NewConnection()  {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Jakarta",
		p.conf.Postges.User,
		p.conf.Postges.Pass,
		p.conf.Postges.Host,
		p.conf.Postges.Port,
		p.conf.Postges.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Error, cannot connect to postgres")
	}

	log.Println("Success connect to PostgreSQL")
	p.DB = db
}

func (p *Postgres) Migration() {
	p.DB.AutoMigrate(&domain.Role{}, &domain.User{}, &domain.Obat{}, &domain.Penyakit{}, &domain.Kunjungan{})
}
