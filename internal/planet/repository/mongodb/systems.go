package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/astro-atlas/internal/planet/models"
	"github.com/Alexander272/astro-atlas/pkg/apperror"
	"github.com/Alexander272/astro-atlas/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SystemsRepo struct {
	db *mongo.Collection
}

func NewSystemRepo(db *mongo.Database, collection string) *SystemsRepo {
	return &SystemsRepo{
		db: db.Collection(collection),
	}
}

func (r *SystemsRepo) Create(ctx context.Context, system models.System) (id string, err error) {
	res, err := r.db.InsertOne(ctx, system)
	if err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convet objectid to hex")
	}
	return oid.Hex(), nil
}

func (r *SystemsRepo) GetList(ctx context.Context) (systems []models.SystemShort, err error) {
	filter := bson.M{}
	cur, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return systems, apperror.ErrNotFound
		}
		return systems, fmt.Errorf("failed to execute query. error: %w", err)
	}
	if err = cur.All(ctx, &systems); err != nil {
		return systems, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return systems, nil
}

func (r *SystemsRepo) GetById(ctx context.Context, id string) (s models.System, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return s, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		logger.Error(res.Err())
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return s, apperror.ErrNotFound
		}
		return s, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err = res.Decode(&s); err != nil {
		return s, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return s, nil
}

func (r *SystemsRepo) Update(ctx context.Context, system models.System) error {
	oid, err := primitive.ObjectIDFromHex(system.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}

	systemByte, err := bson.Marshal(system)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	err = bson.Unmarshal(systemByte, &updateObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal document. error: %w", err)
	}

	delete(updateObj, "_id")
	update := bson.M{"$set": updateObj}

	res, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return apperror.ErrNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *SystemsRepo) Delete(ctx context.Context, systemId string) error {
	oid, err := primitive.ObjectIDFromHex(systemId)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}

	res, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.DeletedCount == 0 {
		return apperror.ErrNotFound
	}

	logger.Tracef("Delete %v documents.\n", res.DeletedCount)
	return nil
}
