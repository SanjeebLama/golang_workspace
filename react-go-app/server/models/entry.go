package models

import (
	"go.mogodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Dish string             `json:"title,omitempty" bson:"title,omitempty"`
	Fat  float64            `json:"content,omitempty" bson:"content,omitempty"`
}
