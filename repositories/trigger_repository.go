package repositories

import (
	"context"
	"workflow_builder/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TriggerRepository struct {
	db *mongo.Collection
}

func NewTriggerRepository(db *mongo.Database) *TriggerRepository {
	return &TriggerRepository{
		db: db.Collection("triggers"),
	}
}

func (repo *TriggerRepository) Create(ctx context.Context, workflow *models.Trigger) error {
	_, err := repo.db.InsertOne(ctx, workflow)
	return err
}
