package client

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/billing"
	billingGrpcGen "github.com/ischenkx/innotech-backend/services/billing/implementation/grpc/pb/generated"
	"google.golang.org/grpc"
)

type Client struct {
	grpcClient billingGrpcGen.BillingClient
}

func (c Client) GetBalances(ctx context.Context, ids []string) ([]float64, error) {
	array, err := c.grpcClient.GetBalances(ctx, &billingGrpcGen.IdArray{Ids: ids})
	if err != nil {
		return nil, err
	}

	return array.Balances, nil
}

func (c Client) Transfer(ctx context.Context, a, b string, amount float64) error {
	_, err := c.grpcClient.Transfer(ctx, &billingGrpcGen.Transaction{
		IdFrom: a,
		IdTo:   b,
		Amount: amount,
	})
	return err
}

func (c Client) GetTransactions(ctx context.Context, dto billing.GetTransactionsDto) ([]billing.Transaction, error) {
	array, err := c.grpcClient.GetTransactions(ctx, &billingGrpcGen.GetTransactionsRequest{
		Id:     dto.Id,
		Offset: dto.Offset,
		Limit:  dto.Limit,
	})
	if err != nil {
		return nil, err
	}

	txs := make([]billing.Transaction, len(array.Transactions))
	for i, tx := range array.Transactions {
		txs[i] = billing.Transaction{
			From:      tx.IdFrom,
			To:        tx.IdTo,
			TimeStamp: tx.Timestamp.AsTime(),
			Amount:    tx.Amount,
		}
	}

	return txs, nil
}

func New(conn grpc.ClientConnInterface) Client {
	return Client{grpcClient: billingGrpcGen.NewBillingClient(conn)}
}
