package client

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/auth"
	auth2 "github.com/ischenkx/innotech-backend/services/auth/implementation/grpc/pb/generated"
	"google.golang.org/grpc"
)

type Client struct {
	grpcClient auth2.AuthClient
}

func (c Client) GenerateTokens(ctx context.Context, data auth.UserData) (auth.Tokens, error) {
	tokens, err := c.grpcClient.GenerateTokens(ctx, &auth2.UserInfo{
		Id:       data.ID,
		Username: data.Username,
	})

	if err != nil {
		return auth.Tokens{}, err
	}

	return auth.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (c Client) Authorize(ctx context.Context, tokens auth.Tokens) (auth.Tokens, auth.UserData, error) {
	res, err := c.grpcClient.Authorize(ctx, &auth2.AuthorizationRequest{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})

	if err != nil {
		return auth.Tokens{}, auth.UserData{}, err
	}

	tokens = auth.Tokens{
		AccessToken:  res.Tokens.AccessToken,
		RefreshToken: res.Tokens.RefreshToken,
	}

	userData := auth.UserData{
		Username: res.UserInfo.Username,
		ID:       res.UserInfo.Id,
	}

	return tokens, userData, nil
}

func New(conn *grpc.ClientConn) Client {
	return Client{auth2.NewAuthClient(conn)}
}

