package db

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/users/service/db/models"
)

// DB is a storage for grabbing
//
// models.User.Username and models.User.Email are unique
type DB interface {
	UpdateUser(ctx context.Context, user models.UpdateUser) (models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUser(ctx context.Context, id string) (models.User, error)
	GetUserByName(ctx context.Context, name string) (models.User, error)
}