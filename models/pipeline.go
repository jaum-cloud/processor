package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pipeline struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Type        string             `bson:"type"`        // "base", "gist", "github"
	CodeSnippet string             `bson:"codeSnippet"` // URL ou identificador do Gist/GitHub
}
