package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/ischenkx/innotech-backend/services/auth"
	"github.com/ischenkx/innotech-backend/services/auth/service/db"
	"github.com/ischenkx/innotech-backend/services/users"
	"log"
	"time"
)

type JWTClaims struct {
	auth.UserData
	jwt.StandardClaims
}

type Service struct {
	users users.Client
	db db.DB
	jwtKey string
	alg string
	tokenTTL time.Duration
}

// Info returns jwt key and alg
func (s *Service) Info() (string, string) {
	return s.jwtKey, s.alg
}

func (s *Service) generateAccessToken(data auth.UserData) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod(s.alg), JWTClaims{UserData: data, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(s.tokenTTL).Unix(),
	}}).
		SignedString([]byte(s.jwtKey))
}

func (s *Service) GenerateTokens(ctx context.Context, data auth.UserData) (auth.Tokens, error) {
	accessToken, err := s.generateAccessToken(data)
	if err != nil {
		return auth.Tokens{}, err
	}
	refreshToken := uuid.New().String()
	if err := s.db.StoreRefreshToken(ctx, refreshToken, data); err != nil {
		return auth.Tokens{}, err
	}
	return auth.Tokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (s *Service) refresh(ctx context.Context, refreshToken string) (auth.Tokens, auth.UserData, error) {

	data, err := s.db.FindRefreshToken(ctx, refreshToken)
	if err != nil {
		return auth.Tokens{}, auth.UserData{}, errors.New("failed to find such refresh token")
	}

	if err := s.db.DeleteRefreshToken(ctx, refreshToken); err != nil {
		return auth.Tokens{}, auth.UserData{}, errors.New("failed to delete refresh token")
	}

	newRefreshToken := uuid.New().String()

	if err := s.db.StoreRefreshToken(ctx, newRefreshToken, data); err != nil {
		return auth.Tokens{}, auth.UserData{}, errors.New("failed to store refresh token")
	}

	accessToken, err := s.generateAccessToken(data)

	if err != nil {
		s.db.DeleteRefreshToken(ctx, newRefreshToken)
		return auth.Tokens{}, auth.UserData{}, fmt.Errorf("failed to generate access token: %s", err)
	}

	return auth.Tokens{
		RefreshToken: newRefreshToken,
		AccessToken:  accessToken,
	}, data, nil
}

func (s *Service) Authorize(ctx context.Context, tokens auth.Tokens) (auth.Tokens, auth.UserData, error) {
	var myClaims JWTClaims
	// don't wanna fuck with those binary operations based errors

	parser := jwt.Parser{SkipClaimsValidation: true}

	_, err := parser.ParseWithClaims(tokens.AccessToken, &myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtKey), nil
	})


	if err != nil {
		if tokens, data, err := s.refresh(ctx, tokens.RefreshToken); err == nil {
			return tokens, data, nil
		} else {
			log.Println("failed to refresh:", err)
		}
		return auth.Tokens{}, auth.UserData{}, err
	}

	if myClaims.ExpiresAt <= time.Now().Unix() {
		tokens, data, err := s.refresh(ctx, tokens.RefreshToken)
		if err != nil {
			return auth.Tokens{}, auth.UserData{}, errors.New("token expired permanently")
		}
		return tokens, data, nil
	}
	return tokens, myClaims.UserData, nil
}

func New(db db.DB, jwtKey string, alg string, expirationTime time.Duration) *Service {
	return &Service{
		db:     db,
		jwtKey: jwtKey,
		alg: alg,
		tokenTTL: expirationTime,
	}
}
