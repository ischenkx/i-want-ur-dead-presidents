package entities

import (
	"context"
)

type EntityInfo struct {
	ID               string
	Title            string
	ShortDesc        string
	LongDesc         string
	MoneyGoal        float64
	OwnerID          string
	DirectorFullName string
	FullCompanyName  string
	Inn              string
	Orgnn            string
	CompanyEmail     string
	OwnerFullName    string
	OwnerPost        string
	PassportData     string
	PictureUrl       string
	ActivityField    string
}

type CreateEntityDto struct {
	Title            string
	ShortDesc        string
	LongDesc         string
	MoneyGoal        float64
	OwnerID          string
	DirectorFullName string
	FullCompanyName  string
	Inn              string
	Orgnn            string
	CompanyEmail     string
	OwnerFullName    string
	OwnerPost        string
	PassportData     string
	PictureUrl       string
	ActivityField    string
}

type DeleteEntityDto struct {
	ID      string
	OwnerID *string
}

type UpdateEntityDto struct {
	ID               string
	Title            *string
	ShortDesc        *string
	LongDesc         *string
	MoneyGoal        *float64
	OwnerID          *string
	DirectorFullName *string
	FullCompanyName  *string
	Inn              *string
	Orgnn            *string
	CompanyEmail     *string
	OwnerFullName    *string
	OwnerPost        *string
	PassportData     *string
	PictureUrl       *string
	ActivityField    *string
}

type GetEntitiesDto struct {
	IDs       []string
	IsPreview bool
}

type GetEntitiesByOwnerIdDto struct {
	OwnerID   string
	IsPreview bool
	Offset    *int64
	Limit     *int64
}

type GetEntitiesRangeDto struct {
	IsPreview bool
	Offset    *int64
	Limit     *int64
}

type Client interface {
	Create(ctx context.Context, dto CreateEntityDto) (EntityInfo, error)
	Delete(ctx context.Context, dto DeleteEntityDto) (EntityInfo, error)
	Update(ctx context.Context, dto UpdateEntityDto) (EntityInfo, error)
	GetRange(ctx context.Context, dto GetEntitiesRangeDto) ([]EntityInfo, error)
	Get(ctx context.Context, dto GetEntitiesDto) ([]EntityInfo, error)
	GetByOwnerID(ctx context.Context, dto GetEntitiesByOwnerIdDto) ([]EntityInfo, error)
}
