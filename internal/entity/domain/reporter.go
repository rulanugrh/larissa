package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reported struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Pengunjung string             `json:"pengunjung" bson:"pengunjung"`
	Age        int             `json:"age" bson:"age"`
	Address    string             `json:"address" bson:"address"`
	CreateAt   time.Time          `json:"create_at" bson:"create_at"`
	UpdateAt   time.Time          `json:"update_at" bson:"update_at"`
}
