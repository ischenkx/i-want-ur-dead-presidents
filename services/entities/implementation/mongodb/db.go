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
	"longdesc": 0,
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
		mutation["ownerid"] = *updateDto.OwnerID
	}
	if updateDto.Title != nil {
		mutation["title"] = *updateDto.Title
	}
	if updateDto.LongDesc != nil {
		mutation["longdesc"] = *updateDto.LongDesc
	}
	if updateDto.ShortDesc != nil {
		mutation["shortdesc"] = *updateDto.ShortDesc
	}
	if updateDto.MoneyGoal != nil {
		mutation["moneygoal"] = *updateDto.MoneyGoal
	}
	if updateDto.DirectorFullName != nil {
		mutation["directorfullname"] = *updateDto.DirectorFullName
	}
	if updateDto.FullCompanyName != nil {
		mutation["fullcompanyname"] = *updateDto.FullCompanyName
	}
	if updateDto.Inn != nil {
		mutation["inn"] = *updateDto.Inn
	}
	if updateDto.Orgnn != nil {
		mutation["orgnn"] = *updateDto.Orgnn
	}
	if updateDto.CompanyEmail != nil {
		mutation["companyemail"] = *updateDto.CompanyEmail
	}
	if updateDto.OwnerFullName != nil {
		mutation["ownerfullname"] = *updateDto.OwnerFullName
	}
	if updateDto.OwnerPost != nil {
		mutation["ownerpost"] = *updateDto.OwnerPost
	}
	if updateDto.PassportData != nil {
		mutation["passportdata"] = *updateDto.PassportData
	}
	if updateDto.PictureUrl != nil {
		mutation["pictureurl"] = *updateDto.PictureUrl
	}
	if updateDto.ActivityField != nil {
		mutation["activityfield"] = *updateDto.ActivityField
	}

	res := db.collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", mutation}}, options.FindOneAndUpdate().SetReturnDocument(options.After))

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
	filter := bson.D{{"ownerid", dto.OwnerID}}

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

func (db *DB) GetRange(ctx context.Context, dto entities.GetEntitiesRangeDto) ([]models.Entity, error) {

	var opts []*options.FindOptions

	if dto.Offset == nil {
		opts = append(opts, options.Find().SetSkip(*dto.Offset))
	}

	if dto.Limit == nil {
		opts = append(opts, options.Find().SetLimit(*dto.Limit))
	}

	res, err := db.collection.Find(ctx, bson.D{}, opts...)

	if err != nil {
		return nil, err
	}

	var ents []models.Entity

	if err := res.Decode(&ents); err != nil {
		return nil, err
	}

	return ents, nil

}

func (db *DB) Delete(ctx context.Context, deleteDto entities.DeleteEntityDto) (models.Entity, error) {
	filter := bson.M{
		"id": deleteDto.ID,
	}

	if deleteDto.OwnerID != nil {
		filter["ownerid"] = *deleteDto.OwnerID
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