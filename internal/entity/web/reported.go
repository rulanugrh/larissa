package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reported struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Pengunjung string             `json:"pengunjung" bson:"pengunjung"`
	Age        int                `json:"age" bson:"age"`
	Address    string             `json:"address" bson:"address"`
	Category   string             `json:"category" bson:"category"`
}
