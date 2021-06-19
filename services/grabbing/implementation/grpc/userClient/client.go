package userClient

import (
	"context"
	"errors"
	grabbing "github.com/ischenkx/innotech-backend/services/grabbing/implementation/grpc/pb/generated"
	"google.golang.org/grpc"
)

type Client struct {
	grpcClient grabbing.GrabbingClient
}

func (s Client) Get(ctx context.Context, product *grabbing.Product) (*grabbing.Response, error) {
	if product == nil {
		return nil, errors.New("nil input")
	}
	response, err := s.grpcClient.Get(ctx, &grabbing.Product{
		Id:  product.Id,
		Inn: product.Inn,
	})
	if err != nil {
		return nil, err
	}
	return &grabbing.Response{
		Score: &grabbing.Score{
			Score:        response.Score.Score,
			CourtScore:   response.Score.CourtScore,
			FinKoefScore: response.Score.FinKoefScore,
			SmartScore:   response.Score.SmartScore,
		},
		Name:     response.Name,
		FullName: response.FullName,
		Inn:      response.Inn,
	}, err
}


func New(client *grpc.ClientConn) Client {
	return Client{grabbing.NewGrabbingClient(client)}
}


