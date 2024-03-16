package repository

import (
	"context"
	"time"

	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportedInterface interface {
	List() (*[]domain.Reported, error)
}

type reported struct {
	client *mongo.Collection
}

func NewReported(client *mongo.Client, conf *config.App) ReportedInterface {
	return &reported{
		client: client.Database(conf.MongoDB.Name).Collection("todo"),
	}
}

func(r *reported) List() (*[]domain.Reported, error) {
	var response []domain.Reported
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	result, err := r.client.Find(ctx, bson.M{})
	if err != nil {
		return nil, util.Errors(err)
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var reported domain.Reported
		if err := result.Decode(&reported); err != nil {
			return nil, util.Errors(err)
		}

		response = append(response, reported)
	}

	return &response, nil
}