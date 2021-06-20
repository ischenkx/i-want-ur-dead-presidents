package service

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/entities"
	"log"

	//"errors" TODO I accidentally deleted half of the code, and now I have no idea where it was used
	"github.com/google/uuid"
	"github.com/ischenkx/innotech-backend/services/entities/service/db"
	"github.com/ischenkx/innotech-backend/services/entities/service/db/models"
)

type Service struct {
	db db.DB
	//validator util.ProductValidator
}

func (s *Service) Delete(ctx context.Context, deleteDto entities.DeleteEntityDto) (entities.EntityInfo, error) {

	log.Println("deleting entity")

	p, err := s.db.Delete(ctx, deleteDto)

	if err != nil {
		return entities.EntityInfo{}, err
	}

	return entities.EntityInfo{
		ID:               p.ID,
		Title:            p.Title,
		ShortDesc:        p.ShortDesc,
		LongDesc:         p.LongDesc,
		MoneyGoal:        p.MoneyGoal,
		OwnerID:          p.OwnerID,
		DirectorFullName: p.DirectorFullName,
		FullCompanyName:  p.FullCompanyName,
		Inn:              p.Inn,
		Orgnn:            p.ORGNN,
		CompanyEmail:     p.CompanyEmail,
		OwnerFullName:    p.OwnerFullName,
		OwnerPost:        p.OwnerPost,
		PassportData:     p.PassportData,
		PictureUrl: p.PictureUrl,
		ActivityField: p.ActivityField,
	}, err
}

func (s *Service) Update(ctx context.Context, updateDto entities.UpdateEntityDto) (entities.EntityInfo, error) {
	//if err := s.validator.ValidateUpdate(updateDto); err != nil {
	//	return entities.EntityInfo{}, err
	//}
	log.Println("updating entity")

	p, err := s.db.Update(ctx, updateDto)

	if err != nil {
		return entities.EntityInfo{}, err
	}

	return entities.EntityInfo{
		ID:               p.ID,
		Title:            p.Title,
		ShortDesc:        p.ShortDesc,
		LongDesc:         p.LongDesc,
		MoneyGoal:        p.MoneyGoal,
		OwnerID:          p.OwnerID,
		DirectorFullName: p.DirectorFullName,
		FullCompanyName:  p.FullCompanyName,
		Inn:              p.Inn,
		Orgnn:            p.ORGNN,
		CompanyEmail:     p.CompanyEmail,
		OwnerFullName:    p.OwnerFullName,
		OwnerPost:        p.OwnerPost,
		PassportData:     p.PassportData,
		PictureUrl: p.PictureUrl,
		ActivityField: p.ActivityField,
	}, err
}

func (s *Service) Create(ctx context.Context, p entities.CreateEntityDto) (entities.EntityInfo, error) {
	//if err := s.validator.ValidateCreate(p); err != nil {
	//	return entities.EntityInfo{}, err
	//}

	log.Println("creating entity")

	mod := models.Entity{
		Title:            p.Title,
		ShortDesc:        p.ShortDesc,
		LongDesc:         p.LongDesc,
		MoneyGoal:        p.MoneyGoal,
		OwnerID:          p.OwnerID,
		DirectorFullName: p.DirectorFullName,
		FullCompanyName:  p.FullCompanyName,
		Inn:              p.Inn,
		ORGNN:            p.Orgnn,
		CompanyEmail:     p.CompanyEmail,
		OwnerFullName:    p.OwnerFullName,
		OwnerPost:        p.OwnerPost,
		PassportData:     p.PassportData,
		PictureUrl: p.PictureUrl,
		ActivityField: p.ActivityField,
	}

	mod.ID = uuid.New().String()

	p2, err := s.db.Create(ctx, mod)

	if err != nil {

		return entities.EntityInfo{}, err
	}
	return entities.EntityInfo{
		ID:               p2.ID,
		OwnerID:          p2.OwnerID,
		Title:            p2.Title,
		ShortDesc:        p2.ShortDesc,
		MoneyGoal:        p2.MoneyGoal,
		LongDesc:         p2.LongDesc,
		DirectorFullName: p2.DirectorFullName,
		FullCompanyName:  p2.FullCompanyName,
		Inn:              p2.Inn,
		Orgnn:            p2.ORGNN,
		CompanyEmail:     p2.CompanyEmail,
		OwnerFullName:    p2.OwnerFullName,
		OwnerPost:        p2.OwnerPost,
		PassportData:     p2.PassportData,
		PictureUrl: p2.PictureUrl,
		ActivityField: p2.ActivityField,
	}, nil

}

func (s *Service) GetRange(ctx context.Context, dto entities.GetEntitiesRangeDto) ([]entities.EntityInfo, error){
	data, err := s.db.GetRange(ctx, dto)
	if err != nil {
		return nil, err
	}

	var infoArray []entities.EntityInfo

	for _, ent := range data {
		infoArray = append(infoArray, entities.EntityInfo{
			ID:               ent.ID,
			Title:            ent.Title,
			ShortDesc:        ent.ShortDesc,
			LongDesc:         ent.LongDesc,
			MoneyGoal:        ent.MoneyGoal,
			OwnerID:          ent.OwnerID,
			DirectorFullName: ent.DirectorFullName,
			FullCompanyName:  ent.FullCompanyName,
			Inn:              ent.Inn,
			Orgnn:            ent.ORGNN,
			CompanyEmail:     ent.CompanyEmail,
			OwnerFullName:    ent.OwnerFullName,
			OwnerPost:        ent.OwnerPost,
			PassportData:     ent.PassportData,
			PictureUrl:       ent.PictureUrl,
			ActivityField:    ent.ActivityField,
		})
	}

	return infoArray, err
}


func (s *Service) Get(ctx context.Context, getEntitiesDto entities.GetEntitiesDto) ([]entities.EntityInfo, error) {
	products, err := s.db.Get(ctx, getEntitiesDto)
	if err != nil {
		return nil, err
	}

	log.Println("retrieving entities")

	var infos []entities.EntityInfo

	for _, p := range products {
		infos = append(infos, entities.EntityInfo{
			ID:               p.ID,
			Title:            p.Title,
			ShortDesc:        p.ShortDesc,
			LongDesc:         p.LongDesc,
			MoneyGoal:        p.MoneyGoal,
			OwnerID:          p.OwnerID,
			DirectorFullName: p.DirectorFullName,
			FullCompanyName:  p.FullCompanyName,
			Inn:              p.Inn,
			Orgnn:            p.ORGNN,
			CompanyEmail:     p.CompanyEmail,
			OwnerFullName:    p.OwnerFullName,
			OwnerPost:        p.OwnerPost,
			PassportData:     p.PassportData,
			PictureUrl: p.PictureUrl,
			ActivityField: p.ActivityField,
		})
	}

	return infos, nil
}

func (s *Service) GetByOwnerID(ctx context.Context, getByOwnerDto entities.GetEntitiesByOwnerIdDto) ([]entities.EntityInfo, error) {
	ents, err := s.db.GetByOwnerID(ctx, getByOwnerDto)

	log.Println("retrieving entities by owner id")

	if err != nil {
		return nil, err
	}

	infos := make([]entities.EntityInfo, 0, len(ents))

	for _, p := range ents {
		infos = append(
			infos, entities.EntityInfo{
				ID:               p.ID,
				Title:            p.Title,
				ShortDesc:        p.ShortDesc,
				LongDesc:         p.LongDesc,
				MoneyGoal:        p.MoneyGoal,
				OwnerID:          p.OwnerID,
				DirectorFullName: p.DirectorFullName,
				FullCompanyName:  p.FullCompanyName,
				Inn:              p.Inn,
				Orgnn:            p.ORGNN,
				CompanyEmail:     p.CompanyEmail,
				OwnerFullName:    p.OwnerFullName,
				OwnerPost:        p.OwnerPost,
				PassportData:     p.PassportData,
				PictureUrl: p.PictureUrl,
				ActivityField: p.ActivityField,
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
