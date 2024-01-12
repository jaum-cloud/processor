package repositories

import (
	"context"
	"workflow_builder/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PipelineRepository struct {
	db *mongo.Collection
}

func NewPipelineRepository(db *mongo.Database) *PipelineRepository {
	return &PipelineRepository{
		db: db.Collection("pipelines"),
	}
}

func (repo *PipelineRepository) Create(ctx context.Context, pipeline *models.Pipeline) error {
	_, err := repo.db.InsertOne(ctx, pipeline)
	return err
}
