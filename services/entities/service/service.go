package service

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/entities"

	//"errors" TODO I accidentally deleted half of the code, and now I have no idea where it was used
	"github.com/google/uuid"
	util "github.com/ischenkx/innotech-backend/services/entities/implementation/util"
	"github.com/ischenkx/innotech-backend/services/entities/service/db"
	"github.com/ischenkx/innotech-backend/services/entities/service/db/models"
)

type Service struct {
	db        db.DB
	validator util.ProductValidator
}

func (s *Service) Delete(ctx context.Context, deleteDto entities.DeleteEntityDto) (entities.EntityInfo, error) {
	p, err := s.db.Delete(ctx, deleteDto)

	if err != nil {
		return entities.EntityInfo{}, err
	}

	return entities.EntityInfo{
		ID:        p.ID,
		Title:     p.Title,
		ShortDesc: p.ShortDesc,
		LongDesc:  p.LongDesc,
		MoneyGoal: p.MoneyGoal,
		OwnerID:   p.OwnerID,
	}, err
}

func (s *Service) Update(ctx context.Context, updateDto entities.UpdateEntityDto) (entities.EntityInfo, error) {
	if err := s.validator.ValidateUpdate(updateDto); err != nil {
		return entities.EntityInfo{}, err
	}

	p, err := s.db.Update(ctx, updateDto)

	if err != nil {
		return entities.EntityInfo{}, err
	}

	return entities.EntityInfo{
		ID:        p.ID,
		Title:     p.Title,
		ShortDesc: p.ShortDesc,
		LongDesc:  p.LongDesc,
		MoneyGoal: p.MoneyGoal,
		OwnerID:   p.OwnerID,
	}, err
}

func (s *Service) Create(ctx context.Context, p entities.CreateEntityDto) (entities.EntityInfo, error) {
	if err := s.validator.ValidateCreate(p); err != nil {
		return entities.EntityInfo{}, err
	}

	mod := models.Entity{
		Title:     p.Title,
		ShortDesc: p.ShortDesc,
		LongDesc:  p.LongDesc,
		MoneyGoal: p.MoneyGoal,
		OwnerID:   p.OwnerID,
	}

	 mod.ID = uuid.New().String()

	p2, err := s.db.Create(ctx, mod)

	if err != nil {

		return entities.EntityInfo{}, err
	}
	return entities.EntityInfo{
		ID:             p2.ID,
		OwnerID:        p2.OwnerID,
		Title:          p2.Title,
		ShortDesc:      p2.ShortDesc,
		MoneyGoal:      p2.MoneyGoal,
		LongDesc:       p2.LongDesc,
	}, nil

}

func (s *Service) Get(ctx context.Context, getEntitiesDto entities.GetEntitiesDto) ([]entities.EntityInfo, error) {
	products, err := s.db.Get(ctx, getEntitiesDto)
	if err != nil {
		return nil, err
	}

	var infos []entities.EntityInfo

	for _, p := range products {
		infos = append(infos, entities.EntityInfo{
			ID:        p.ID,
			Title:     p.Title,
			ShortDesc: p.ShortDesc,
			LongDesc:  p.LongDesc,
			MoneyGoal: p.MoneyGoal,
			OwnerID:   p.OwnerID,
		})
	}

	return infos, nil
}

func (s *Service) GetByOwnerID(ctx context.Context, getByOwnerDto entities.GetEntitiesByOwnerIdDto) ([]entities.EntityInfo, error) {
	ents, err := s.db.GetByOwnerID(ctx, getByOwnerDto)

	if err != nil {
		return nil, err
	}

	infos := make([]entities.EntityInfo, 0, len(ents))

	for _, p := range ents {
		infos = append(
			infos, entities.EntityInfo{
				ID:        p.ID,
				Title:     p.Title,
				ShortDesc: p.ShortDesc,
				LongDesc:  p.LongDesc,
				MoneyGoal: p.MoneyGoal,
				OwnerID:   p.OwnerID,
			},
		)
	}

	return infos, nil
}

func New(db db.DB) *Service {
	return &Service{
		db: db,
	}
}
