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

type PlanetRepo struct {
	db *mongo.Collection
}

func NewPlanetRepo(db *mongo.Database, collection string) *PlanetRepo {
	return &PlanetRepo{
		db: db.Collection(collection),
	}
}

func (r *PlanetRepo) Create(ctx context.Context, planet models.Planet) (id string, err error) {
	res, err := r.db.InsertOne(ctx, planet)
	if err != nil {
		return "", fmt.Errorf("failed to execute query. error: %w", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convet objectid to hex")
	}
	return oid.Hex(), nil
}

func (r *PlanetRepo) GetList(ctx context.Context, systemId string) (planets []models.PlanetShort, err error) {
	filter := bson.M{"systemId": systemId}
	cur, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return planets, apperror.ErrNotFound
		}
		return planets, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err = cur.All(ctx, &planets); err != nil {
		return planets, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return planets, nil
}

func (r *PlanetRepo) GetById(ctx context.Context, planetId string) (p models.Planet, err error) {
	oid, err := primitive.ObjectIDFromHex(planetId)
	if err != nil {
		return p, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		logger.Error(res.Err())
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return p, apperror.ErrNotFound
		}
		return p, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}

	if err = res.Decode(&p); err != nil {
		return p, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return p, nil
}

func (r *PlanetRepo) Update(ctx context.Context, planet models.Planet) error {
	oid, err := primitive.ObjectIDFromHex(planet.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	planetByte, err := bson.Marshal(planet)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	err = bson.Unmarshal(planetByte, updateObj)
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

func (r *PlanetRepo) Delete(ctx context.Context, planetId string) error {
	oid, err := primitive.ObjectIDFromHex(planetId)
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
