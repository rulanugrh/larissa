package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Conn *mongo.Client
	conf *App
}

func InitializeMongo(conf *App) *MongoDB {
	return &MongoDB{conf: conf}
}

func (m *MongoDB) NewMongo() {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	dsn := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		m.conf.MongoDB.User,
		m.conf.MongoDB.Pass,
		m.conf.MongoDB.Host,
	)

	client := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverApi)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cli, err := mongo.Connect(ctx, client)
	if err != nil {
		log.Println("Error, can't connect mongodb")
	}

	log.Println("Success connect to MongoDB")
	m.Conn = cli
}
