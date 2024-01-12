package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workflow struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Trigger   Trigger            `bson:"trigger"`
	Pipelines []Pipeline         `bson:"pipelines"`
}
