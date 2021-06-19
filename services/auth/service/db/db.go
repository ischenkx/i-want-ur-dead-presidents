package db

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/auth"
)

// DB is a storage for grabbing
//
// models.User.Username and models.User.Email are unique
type DB interface {
	StoreRefreshToken(context.Context, string, auth.UserData) error
	DeleteRefreshToken(context.Context, string) error
	FindRefreshToken(context.Context, string) (auth.UserData, error)
}