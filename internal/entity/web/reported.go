package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reported struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Pengunjung string             `json:"pengunjung" bson:"pengunjung"`
	Age        string             `json:"age" bson:"age"`
	Address    string             `json:"address" bson:"address"`
}