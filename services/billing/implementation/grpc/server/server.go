package server

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/billing"
	billingGrpcGen "github.com/ischenkx/innotech-backend/services/billing/implementation/grpc/pb/generated"
	"github.com/ischenkx/innotech-backend/services/billing/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	billingGrpcGen.UnimplementedBillingServer
	service *service.Service
}

func (s *Server) GetBalances(ctx context.Context, req *billingGrpcGen.IdArray) (*billingGrpcGen.BalanceArray, error) {
	balances, err := s.service.GetBalances(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}

	var arr billingGrpcGen.BalanceArray

	for _, balance := range balances {
		arr.Balances = append(arr.Balances, balance)
	}

	return &arr, nil
}
func (s *Server) GetTransactions(ctx context.Context, req *billingGrpcGen.GetTransactionsRequest) (*billingGrpcGen.TransactionArray, error) {
	txs, err := s.service.GetTransactions(ctx, billing.GetTransactionsDto{
		Offset: req.Offset,
		Limit:  req.Limit,
		Id:     req.Id,
	})

	if err != nil {
		return nil, err
	}

	var arr billingGrpcGen.TransactionArray

	for _, tx := range txs {
		arr.Transactions = append(arr.Transactions, &billingGrpcGen.Transaction{
			IdFrom:    tx.From,
			IdTo:      tx.To,
			Amount:    tx.Amount,
			Timestamp: timestamppb.New(tx.TimeStamp),
		})
	}

	return &arr, nil
}
func (s *Server) Transfer(ctx context.Context, req *billingGrpcGen.Transaction) (*billingGrpcGen.Empty, error) {
	return &billingGrpcGen.Empty{}, s.service.Transfer(ctx, req.IdFrom, req.IdTo, req.Amount)
}

func New(srv *service.Service) *Server {
	return &Server{
		service: srv,
	}
}