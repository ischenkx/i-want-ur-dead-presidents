package mongodb

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenRecord struct {
	Token string
	Data auth.UserData
}

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

func(db *DB) StoreRefreshToken(ctx context.Context, token string, data auth.UserData) error {
	_, err := db.collection.InsertOne(ctx, TokenRecord{
		Token: token,
		Data:  data,
	})
	return err
}

func (db *DB) DeleteRefreshToken(ctx context.Context, token string) error {
	_, err := db.collection.DeleteOne(ctx, bson.D{{"token", token}})
	return err
}

func (db *DB) FindRefreshToken(ctx context.Context, token string) (auth.UserData, error) {
	res := db.collection.FindOne(ctx, bson.D{{"token", token}})
	if res.Err() != nil {
		return auth.UserData{}, res.Err()
	}
	var rec TokenRecord
	if err := res.Decode(&rec); err != nil {
		return auth.UserData{}, err
	}
	return rec.Data, nil
}

func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}