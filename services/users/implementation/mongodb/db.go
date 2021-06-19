package mongodb

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/users/service/db/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (db *DB) UpdateUser(ctx context.Context, updateDto models.UpdateUser) (models.User, error) {

	mutation := bson.M{}

	if updateDto.Username != nil {
		mutation["username"] = *updateDto.Username
	}
	if updateDto.Password != nil {
		mutation["password"] = *updateDto.Password
	}

	res := db.collection.FindOneAndUpdate(ctx,
		bson.D{{"id", updateDto.ID}},
		bson.D{
			{
				"$set",
				mutation,
			},
		},
		options.FindOneAndUpdate().
			SetReturnDocument(options.After),
		)

	if res.Err() != nil {
		return models.User{}, res.Err()
	}

	var u models.User

	if err := res.Decode(&u); err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (db *DB) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	_, err := db.collection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (db *DB) GetUser(ctx context.Context, id string) (models.User, error) {
	filter := bson.D{{"id", id}}
	result := db.collection.FindOne(ctx, filter)

	var user models.User
	err := result.Decode(&user)
	return user, err
}

func (db *DB) GetUserByName(ctx context.Context, username string) (models.User, error) {
	filter := bson.D{{"username", username}}
	result := db.collection.FindOne(ctx, filter)

	var user models.User
	err := result.Decode(&user)
	return user, err
}

func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}