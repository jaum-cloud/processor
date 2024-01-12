package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trigger struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Type  string             `bson:"type"` // "schedule", "manual", "webhook"
	Value string             `bson:"value"`
}
