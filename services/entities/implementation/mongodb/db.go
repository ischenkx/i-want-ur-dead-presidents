package mongodb

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/entities"
	"github.com/ischenkx/innotech-backend/services/entities/service/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var previewProjection = bson.M{
	"long_desc": 0,
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

func (db *DB) Update(ctx context.Context, updateDto entities.UpdateEntityDto) (models.Entity, error) {

	filter := bson.M{
		"id": updateDto.ID,
	}

	mutation := bson.M{}

	if updateDto.OwnerID != nil {
		mutation["owner_id"] = *updateDto.OwnerID
	}
	if updateDto.Title != nil {
		mutation["title"] = *updateDto.Title
	}
	if updateDto.LongDesc != nil {
		mutation["long_desc"] = *updateDto.LongDesc
	}
	if updateDto.ShortDesc != nil {
		mutation["short_desc"] = *updateDto.ShortDesc
	}
	if updateDto.MoneyGoal != nil {
		mutation["money_goal"] = *updateDto.MoneyGoal
	}

	res := db.collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", mutation}})

	if res.Err() != nil {
		return models.Entity{}, res.Err()
	}

	var result models.Entity

	if err := res.Decode(&result); err != nil {
		return models.Entity{}, err
	}

	return result, nil
}

func (db *DB) Create(ctx context.Context, p models.Entity) (models.Entity, error) {
	_, err := db.collection.InsertOne(ctx, p)
	if err != nil {
		return models.Entity{}, err
	}
	return p, err
}

func (db *DB) Get(ctx context.Context, dto entities.GetEntitiesDto) ([]models.Entity, error) {

	var findOpts []*options.FindOptions

	if dto.IsPreview {
		findOpts = append(findOpts, options.Find().SetProjection(previewProjection))
	}

	filter := bson.D{
			{"id", bson.D{
				{"$in", dto.IDs },
			},
		},
	}
	cursor, err := db.collection.Find(ctx, filter, findOpts...)
	if err != nil {
		return nil, err
	}
	var ents []models.Entity
	if err := cursor.All(ctx, &ents); err != nil {
		return nil, err
	}
	return ents, err
}

func (db *DB) GetByOwnerID(ctx context.Context, dto entities.GetEntitiesByOwnerIdDto) ([]models.Entity, error) {
	filter := bson.D{{"owner_id", dto.OwnerID}}

	var findOpts []*options.FindOptions

	if dto.IsPreview {
		findOpts = append(findOpts, options.Find().SetProjection(previewProjection))
	}

	if dto.Offset != nil {
		if *dto.Offset > 0 {
			findOpts = append(findOpts, options.Find().SetSkip(*dto.Offset))
		}
	}


	if dto.Limit != nil {
		if *dto.Limit > 0 {
			findOpts = append(findOpts, options.Find().SetLimit(*dto.Limit))
		}
	}

	result, err := db.collection.Find(ctx, filter, findOpts...)
	if err != nil {
		return nil, err
	}

	var ents []models.Entity
	err = result.All(ctx, &ents)
	return ents, err
}

func (db *DB) Delete(ctx context.Context, deleteDto entities.DeleteEntityDto) (models.Entity, error) {
	filter := bson.M{
		"id": deleteDto.ID,
	}

	if deleteDto.OwnerID != nil {
		filter["owner_id"] = *deleteDto.OwnerID
	}

	res := db.collection.FindOneAndDelete(ctx, filter)

	if res.Err() != nil {
		return models.Entity{}, res.Err()
	}

	var ent models.Entity
	if err := res.Decode(&ent); err != nil {
		return models.Entity{}, err
	}
	return ent, nil
}


func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}