package mongodb

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/billing"
	"github.com/ischenkx/innotech-backend/services/billing/service/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DB struct {
	client     *mongo.Client
	wallets *mongo.Collection
	transactions *mongo.Collection
}

func (db *DB) Transfer(ctx context.Context, from, to string, amount float64) error {
	sess, err := db.client.StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx)

	_, err = sess.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

		if from != "" {
			res := db.wallets.FindOneAndUpdate(sessCtx, bson.M{
				"id": from,
				"balance": bson.M{"$gte": amount},
			}, bson.D{{"$inc", bson.D{
				{"balance", -amount},
			}}})

			if res.Err() != nil {
				sessCtx.AbortTransaction(ctx)
				return nil, err
			}
		}



		res := db.wallets.FindOneAndUpdate(sessCtx, bson.M{"id": to,}, bson.D{{"$inc", bson.D{
			{"balance", amount},
		}}}, options.FindOneAndUpdate().SetUpsert(true))

		if res.Err() != nil {
			sessCtx.AbortTransaction(ctx)
			return nil, err
		}

		_, err := db.transactions.InsertOne(sessCtx, billing.Transaction{
			From:      from,
			To:        to,
			TimeStamp: time.Now(),
			Amount:    amount,
		})

		if err != nil {
			sessCtx.AbortTransaction(ctx)
			return nil, err
		}

		sessCtx.CommitTransaction(ctx)

		return nil, nil
	})

	return err
}

func (db *DB) GetTransactions(ctx context.Context, dto billing.GetTransactionsDto) ([]billing.Transaction, error) {
	var findOpts []*options.FindOptions

	filter := bson.D{
		{"from", dto.Id },
		{"to", dto.Id },
	}
	cursor, err := db.transactions.Find(ctx, filter, findOpts...)
	if err != nil {
		return nil, err
	}
	var transactions []billing.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}
	return transactions, err
}

func (db *DB) GetBalances(ctx context.Context, ids []string) ([]float64, error) {
	var findOpts []*options.FindOptions

	filter := bson.D{
		{"id", bson.D{
			{"$in", ids},
		},
		},
	}
	cursor, err := db.wallets.Find(ctx, filter, findOpts...)
	if err != nil {
		return nil, err
	}
	var wallets []models.Wallet
	if err := cursor.All(ctx, &wallets); err != nil {
		return nil, err
	}

	balances := make([]float64, len(wallets))
	for i, wallet := range wallets {
		balances[i] = wallet.Balance
	}
	return balances, err
}

func Connect(url, db, walletsCollection, transactionsCollection string) (*DB, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	return &DB{
		client:     client,
		transactions: client.Database(db).Collection(transactionsCollection),
		wallets: client.Database(db).Collection(walletsCollection),
	}, nil
}

func (db *DB) Close(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}
