package server

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/auth"
	auth2 "github.com/ischenkx/innotech-backend/services/auth/implementation/grpc/pb/generated"
	"github.com/ischenkx/innotech-backend/services/auth/service"
)

type Server struct {
	auth2.UnimplementedAuthServer
	service *service.Service
}


func (s *Server) GenerateTokens(ctx context.Context, info *auth2.UserInfo) (*auth2.TokensReply, error) {
	tokens, err := s.service.GenerateTokens(ctx, auth.UserData{
		Username: info.Username,
		ID:       info.Id,
	})

	if err != nil {
		return nil, err
	}

	return &auth2.TokensReply{
		RefreshToken: tokens.RefreshToken,
		AccessToken:  tokens.AccessToken,
	}, nil
}

func (s *Server) Authorize(ctx context.Context, req *auth2.AuthorizationRequest) (*auth2.AuthorizationReply, error) {
	tokens, data, err := s.service.Authorize(ctx, auth.Tokens{
		RefreshToken: req.RefreshToken,
		AccessToken:  req.AccessToken,
	})

	if err != nil {
		return nil, err
	}

	return &auth2.AuthorizationReply{
		UserInfo: &auth2.UserInfo{
			Id:       data.ID,
			Username: data.Username,
		},
		Tokens:   &auth2.TokensReply{
			RefreshToken: tokens.RefreshToken,
			AccessToken:  tokens.AccessToken,
		},
	}, nil
}

func (s *Server) GetInfo(ctx context.Context, req *auth2.EmptyRequest) (*auth2.JWTInfo, error) {
	key, alg := s.service.Info()

	return &auth2.JWTInfo{
		Alg: alg,
		Key: key,
	}, nil
}

func New(srv *service.Service) *Server {
	return &Server{
		service: srv,
	}
}