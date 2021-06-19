package auth

import "context"

type Tokens struct {
	AccessToken string
	RefreshToken string
}

type UserData struct {
	Username, ID string
}

type Client interface {
	GenerateTokens(ctx context.Context, data UserData) (Tokens, error)
	Authorize(ctx context.Context, tokens Tokens) (Tokens, UserData, error)
}
