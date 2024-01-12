package repositories

import (
	"context"
	"workflow_builder/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type WorkflowRepository struct {
	db *mongo.Collection
}

func NewWorkflowRepository(db *mongo.Database) *WorkflowRepository {
	return &WorkflowRepository{
		db: db.Collection("workflows"),
	}
}

func (repo *WorkflowRepository) Create(ctx context.Context, workflow *models.Workflow) error {
	_, err := repo.db.InsertOne(ctx, workflow)
	return err
}
