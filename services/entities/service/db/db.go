package db

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/entities"
	"github.com/ischenkx/innotech-backend/services/entities/service/db/models"
)

// DB is a storage for Products
//
// models.Product.ID is unique
type DB interface {
	Update(ctx context.Context, updateDto entities.UpdateEntityDto) (models.Entity, error)
	Create(ctx context.Context, ent models.Entity) (models.Entity, error)
	Get(ctx context.Context, getDto entities.GetEntitiesDto) ([]models.Entity, error)
	GetByOwnerID(ctx context.Context, getByOwnerIdDto entities.GetEntitiesByOwnerIdDto) ([]models.Entity, error)
	Delete(ctx context.Context, deleteDto entities.DeleteEntityDto) (models.Entity, error)
}