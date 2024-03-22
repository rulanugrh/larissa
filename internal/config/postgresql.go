package config

import (
	"fmt"
	"log"
	"time"

	"github.com/rulanugrh/larissa/internal/entity/domain"
	"golang.org/x/crypto/bcrypt"
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
	err := p.DB.AutoMigrate(&domain.Role{}, &domain.User{}, &domain.Obat{}, &domain.Penyakit{}, &domain.Kunjungan{})
	if err != nil {
		log.Printf("error migration because :%s", err.Error())
	}

	log.Println("success migration")
}

func (p *Postgres) Seeder() error {
	roles := append([]domain.Role{},
		domain.Role{
			Name: "Admin",
			Descriptions: "this is role admin",
		},
		domain.Role{
			Name: "Doctor",
			Descriptions: "this is role doctor",
		},
		domain.Role{
			Name: "Nurse",
			Descriptions: "this is role nurse",
		},
		domain.Role{
			Name: "User",
			Descriptions: "this is role for user",
	})

	err := p.DB.Create(&roles).Error
	if err != nil {
		return err
	}

	password, err:= bcrypt.GenerateFromPassword([]byte(p.conf.Admin.Password), 14)
	if err != nil {
		return err
	}

	user := domain.User{
		FName: "Admini",
		LName: "-",
		Email: p.conf.Admin.Email,
		Password: string(password),
		Address: "-",
		Age: 0,
		RoleID: 1,
		TTL: time.Now(),
	}

	err = p.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
