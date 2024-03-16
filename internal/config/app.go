package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Postges struct {
		Name string
		Host string
		Port string
		User string
		Pass string
	}

	MongoDB struct {
		Host string
		Pass string
		User string
	}

	App struct {
		Secret string
		Host   string
		Port   string
	}

	Admin struct {
		Email    string
		Password string
	}
}
var app *App

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}
	
	return app
}

// Initialize all config in .env
func initConfig() *App {
	conf := App{}
	err := godotenv.Load()
	if err != nil {
		conf.Postges.User = "root"
		conf.Postges.Port = "5432"
		conf.Postges.Pass = ""
		conf.Postges.Name = "larissa.db"
		conf.Postges.Host = "localhost"

		conf.Admin.Email = ""
		conf.Admin.Password = ""

		conf.MongoDB.Host = "localhost"
		conf.MongoDB.Pass = "larissa"
		conf.MongoDB.User = "larissa"

		conf.App.Host = "localhost"
		conf.App.Port = "3000"
		conf.App.Secret = "s3cret"

		return &conf
	}

	conf.Postges.Host = os.Getenv("POSTGRES_HOST")
	conf.Postges.Port = os.Getenv("POSTGRES_PORT")
	conf.Postges.User = os.Getenv("POSTGRES_USER")
	conf.Postges.Pass = os.Getenv("POSTGRES_PASS")
	conf.Postges.Name = os.Getenv("POSTGRES_NAME")

	conf.Admin.Email = os.Getenv("ADMIN_EMAIL")
	conf.Admin.Password = os.Getenv("ADMIN_PASSWORD")

	conf.MongoDB.Host = os.Getenv("MONGODB_HOST")
	conf.MongoDB.Pass = os.Getenv("MONGODB_PASS")
	conf.MongoDB.User = os.Getenv("MONGODB_USER")

	conf.App.Secret = os.Getenv("APP_SECRET")
	conf.App.Port = os.Getenv("APP_PORT")
	conf.App.Host = os.Getenv("APP_HOST")

	return &conf
}