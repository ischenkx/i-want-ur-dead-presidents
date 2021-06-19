package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
	collection *mongo.Collection
}

func Connect(url, db, collection string) (*DB, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return &DB{
		client: client,
		collection: client.Database(db).Collection(collection),
	}, nil
}

func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}