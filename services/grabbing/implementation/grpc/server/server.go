package server

import (
	"context"
	"errors"
	"github.com/ischenkx/innotech-backend/services/grabbing/implementation/grpc/pb/generated"
	"github.com/ischenkx/innotech-backend/services/grabbing/service"
	"github.com/ischenkx/innotech-backend/services/grabbing/service/db/models"
)

type Server struct {
	grabbing.UnimplementedGrabbingServer
	service *service.Service
}

func (s *Server) Get(ctx context.Context, product *grabbing.Product) (*grabbing.Response, error) {
	if product == nil {
		return nil, errors.New("nil input")
	}
	response, err := s.service.Get(ctx, models.Product{
		Id:  product.Id,
		Inn: product.Inn,
	})
	if err != nil {
		return nil, err
	}
	return &grabbing.Response{
		Score: &grabbing.Score{
			Score:        int32(response.Score.OverallScore),
			CourtScore:   int32(response.Score.CourtScore),
			FinKoefScore: int32(response.Score.FinKoefScore),
			SmartScore:   int32(response.Score.SmartScore),
		},
		Name:     response.Name,
		FullName: response.FullName,
		Inn:      response.Inn,
	}, err
}

func New(srv *service.Service) *Server {
	return &Server{
		service: srv,
	}
}
