package billing

import (
	"context"
	"time"
)

type Transaction struct {
	From string
	To string
	TimeStamp time.Time

	Amount float64
}

type GetTransactionsDto struct {
	Offset *int64
	Limit *int64
	Id string
}

type Client interface {
	GetBalances(ctx context.Context, ids []string) ([]float64, error)
	Transfer(ctx context.Context, a, b string, amount float64) error
	GetTransactions(ctx context.Context, dto GetTransactionsDto) ([]Transaction, error)
}
