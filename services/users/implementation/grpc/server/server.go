package server

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/users/implementation/grpc/pb/generated"
	"github.com/ischenkx/innotech-backend/services/users/service"
)

var emptyReply = &users.EmptyReply{}

type Server struct {
	users.UnimplementedUsersServer
	service *service.Service
}

func (s *Server) UpdateUsername(ctx context.Context, req *users.UpdateUsernameRequest) (*users.EmptyReply, error) {
	err := s.service.UpdateUsername(ctx, req.Id, req.Username)
	return emptyReply, err
}

func (s *Server) UpdatePassword(ctx context.Context, req *users.UpdatePasswordRequest) (*users.EmptyReply, error) {
	err := s.service.UpdatePassword(ctx, req.Id, req.PreviousPassword, req.Password)
	return emptyReply, err
}

func (s *Server) Login(ctx context.Context, req *users.LoginRequest) (*users.User, error) {
	u, err := s.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &users.User{
		Username: u.Username,
		Id:       u.ID,
	}, nil
}

func (s *Server) Register(ctx context.Context, req *users.RegisterRequest) (*users.User, error) {
	u, err := s.service.Register(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &users.User{
		Username: u.Username,
		Id:       u.ID,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *users.GetUserRequest) (*users.User, error) {
	u, err := s.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &users.User{
		Username: u.Username,
		Id:       u.ID,
	}, nil
}

func (s *Server) GetByName(ctx context.Context, req *users.GetUserByNameRequest) (*users.User, error) {
	u, err := s.service.GetByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &users.User{
		Username: u.Username,
		Id:       u.ID,
	}, nil
}

func New(srv *service.Service) *Server {
	return &Server{
		service: srv,
	}
}