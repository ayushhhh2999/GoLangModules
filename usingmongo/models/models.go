package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Course represents a course in the watchlist
type Netflix struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Movie   string             `bson:"title,omitempty" json:"title,omitempty"`
	Watched bool               `bson:"watched,omitempty" json:"watched,omitempty"`
}
