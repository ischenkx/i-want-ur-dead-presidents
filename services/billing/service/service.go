package service

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/billing"
	"github.com/ischenkx/innotech-backend/services/billing/service/db"
)

type Service struct {
	db db.DB
}

func (s *Service) GetBalances(ctx context.Context, ids []string) ([]float64, error) {
	return s.db.GetBalances(ctx, ids)
}

func (s *Service) Transfer(ctx context.Context, a, b string, amount float64) error {
	return s.db.Transfer(ctx, a, b, amount)
}

func (s *Service) GetTransactions(ctx context.Context, dto billing.GetTransactionsDto) ([]billing.Transaction, error) {
	return s.db.GetTransactions(ctx, dto)
}

func New(db db.DB) *Service {
	return &Service{db}
}