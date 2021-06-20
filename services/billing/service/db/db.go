package db

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/billing"
)

type DB interface {
	Transfer(ctx context.Context, from, to string, amount float64) error
	GetTransactions(ctx context.Context, dto billing.GetTransactionsDto) ([]billing.Transaction, error)
	GetBalances(ctx context.Context, ids []string) ([]float64, error)
}