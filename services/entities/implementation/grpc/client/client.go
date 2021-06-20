package client

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/entities"
	entitiesGrpcGen "github.com/ischenkx/innotech-backend/services/entities/implementation/grpc/pb/generated"
	"google.golang.org/grpc"
)

type Client struct {
	grpcClient entitiesGrpcGen.EntitiesClient
}

func (c Client) Create(ctx context.Context, dto entities.CreateEntityDto) (entities.EntityInfo, error) {
	info, err := c.grpcClient.Create(ctx, &entitiesGrpcGen.CreateEntityRequest{
		Title:            dto.Title,
		ShortDesc:        dto.ShortDesc,
		LongDesc:         dto.LongDesc,
		MoneyGoal:        dto.MoneyGoal,
		OwnerId:          dto.OwnerID,
		DirectorFullName: dto.DirectorFullName,
		FullCompanyName:  dto.FullCompanyName,
		Inn:              dto.Inn,
		Orgnn:            dto.Orgnn,
		CompanyEmail:     dto.CompanyEmail,
		OwnerFullName:    dto.OwnerFullName,
		OwnerPost:        dto.OwnerPost,
		PassportData:     dto.PassportData,
		ActivityField:    dto.ActivityField,
	})

	if err != nil {
		return entities.EntityInfo{}, err
	}

	return entities.EntityInfo{
		ID:               info.Id,
		Title:            info.Title,
		ShortDesc:        info.ShortDesc,
		LongDesc:         info.LongDesc,
		MoneyGoal:        info.MoneyGoal,
		OwnerID:          info.OwnerId,
		DirectorFullName: info.DirectorFullName,
		FullCompanyName:  info.FullCompanyName,
		Inn:              info.Inn,
		Orgnn:            info.Orgnn,
		CompanyEmail:     info.CompanyEmail,
		OwnerFullName:    info.OwnerFullName,
		OwnerPost:        info.OwnerPost,
		PassportData:     info.PassportData,
		//PictureUrl:       info.PictureUrl,
		ActivityField:    info.ActivityField,
	}, nil
}

func (c Client) Delete(ctx context.Context, dto entities.DeleteEntityDto) (entities.EntityInfo, error) {
	info, err := c.grpcClient.Delete(ctx, &entitiesGrpcGen.DeleteEntityRequest{
		Id:      dto.ID,
		OwnerId: dto.OwnerID,
	})

	if err != nil {
		return entities.EntityInfo{}, err
	}

	return entities.EntityInfo{
		ID:               info.Id,
		Title:            info.Title,
		ShortDesc:        info.ShortDesc,
		LongDesc:         info.LongDesc,
		MoneyGoal:        info.MoneyGoal,
		OwnerID:          info.OwnerId,
		DirectorFullName: info.DirectorFullName,
		FullCompanyName:  info.FullCompanyName,
		Inn:              info.Inn,
		Orgnn:            info.Orgnn,
		CompanyEmail:     info.CompanyEmail,
		OwnerFullName:    info.OwnerFullName,
		OwnerPost:        info.OwnerPost,
		PassportData:     info.PassportData,
		//PictureUrl:       info.PictureUrl,
		ActivityField:    info.ActivityField,
	}, nil
}

func (c Client) Update(ctx context.Context, dto entities.UpdateEntityDto) (entities.EntityInfo, error) {
	info, err := c.grpcClient.Update(ctx, &entitiesGrpcGen.UpdateEntityRequest{
		Title:            dto.Title,
		ShortDesc:        dto.ShortDesc,
		LongDesc:         dto.LongDesc,
		MoneyGoal:        dto.MoneyGoal,
		OwnerId:          dto.OwnerID,
		Id:               dto.ID,
		DirectorFullName: dto.DirectorFullName,
		FullCompanyName:  dto.FullCompanyName,
		Inn:              dto.Inn,
		Orgnn:            dto.Orgnn,
		CompanyEmail:     dto.CompanyEmail,
		OwnerFullName:    dto.OwnerFullName,
		OwnerPost:        dto.OwnerPost,
		PassportData:     dto.PassportData,
		//PictureUrl:       dto.PictureUrl,
		ActivityField:    dto.ActivityField,
	})

	if err != nil {
		return entities.EntityInfo{}, err
	}

	return entities.EntityInfo{
		ID:               info.Id,
		Title:            info.Title,
		ShortDesc:        info.ShortDesc,
		LongDesc:         info.LongDesc,
		MoneyGoal:        info.MoneyGoal,
		OwnerID:          info.OwnerId,
		DirectorFullName: info.DirectorFullName,
		FullCompanyName:  info.FullCompanyName,
		Inn:              info.Inn,
		Orgnn:            info.Orgnn,
		CompanyEmail:     info.CompanyEmail,
		OwnerFullName:    info.OwnerFullName,
		OwnerPost:        info.OwnerPost,
		PassportData:     info.PassportData,
		//PictureUrl:       info.PictureUrl,
		ActivityField:    info.ActivityField,
	}, nil
}

func (c Client) Get(ctx context.Context, dto entities.GetEntitiesDto) ([]entities.EntityInfo, error) {
	ents, err := c.grpcClient.Get(ctx, &entitiesGrpcGen.GetEntityRequest{
		Ids: &entitiesGrpcGen.IDArray{
			Ids: dto.IDs,
		},
		IsPreview: dto.IsPreview,
	})

	if err != nil {
		return nil, err
	}

	var res []entities.EntityInfo

	for _, e := range ents.Entities {
		res = append(res, entities.EntityInfo{
			ID:               e.Id,
			Title:            e.Title,
			ShortDesc:        e.ShortDesc,
			LongDesc:         e.LongDesc,
			MoneyGoal:        e.MoneyGoal,
			OwnerID:          e.OwnerId,
			DirectorFullName: e.DirectorFullName,
			FullCompanyName:  e.FullCompanyName,
			Inn:              e.Inn,
			Orgnn:            e.Orgnn,
			CompanyEmail:     e.CompanyEmail,
			OwnerFullName:    e.OwnerFullName,
			OwnerPost:        e.OwnerPost,
			PassportData:     e.PassportData,
			//PictureUrl:       e.PictureUrl,
			ActivityField:    e.ActivityField,
		})
	}

	return res, nil
}

func (c Client) GetRange(ctx context.Context, dto entities.GetEntitiesRangeDto) ([]entities.EntityInfo, error) {

	ents, err := c.grpcClient.GetRange(ctx, &entitiesGrpcGen.GetRangeEntityRequest{
		Offset:    dto.Offset,
		Limit:     dto.Limit,
		IsPreview: dto.IsPreview,
	})

	if err != nil {
		return nil, err
	}

	var res []entities.EntityInfo

	for _, e := range ents.Entities {
		res = append(res, entities.EntityInfo{
			ID:               e.Id,
			Title:            e.Title,
			ShortDesc:        e.ShortDesc,
			LongDesc:         e.LongDesc,
			MoneyGoal:        e.MoneyGoal,
			OwnerID:          e.OwnerId,
			DirectorFullName: e.DirectorFullName,
			FullCompanyName:  e.FullCompanyName,
			Inn:              e.Inn,
			Orgnn:            e.Orgnn,
			CompanyEmail:     e.CompanyEmail,
			OwnerFullName:    e.OwnerFullName,
			OwnerPost:        e.OwnerPost,
			PassportData:     e.PassportData,
			//PictureUrl:       e.PictureUrl,
			ActivityField:    e.ActivityField,
		})
	}

	return res, nil
}

func (c Client) GetByOwnerID(ctx context.Context, dto entities.GetEntitiesByOwnerIdDto) ([]entities.EntityInfo, error) {
	ents, err := c.grpcClient.GetByOwnerID(ctx, &entitiesGrpcGen.GetEntityByOwnerIDRequest{
		OwnerId:   dto.OwnerID,
		Offset:    dto.Offset,
		Limit:     dto.Limit,
		IsPreview: dto.IsPreview,
	})
	if err != nil {
		return nil, err
	}
	var res []entities.EntityInfo
	for _, e := range ents.Entities {
		res = append(res, entities.EntityInfo{
			ID:               e.Id,
			Title:            e.Title,
			ShortDesc:        e.ShortDesc,
			LongDesc:         e.LongDesc,
			MoneyGoal:        e.MoneyGoal,
			OwnerID:          e.OwnerId,
			DirectorFullName: e.DirectorFullName,
			FullCompanyName:  e.FullCompanyName,
			Inn:              e.Inn,
			Orgnn:            e.Orgnn,
			CompanyEmail:     e.CompanyEmail,
			OwnerFullName:    e.OwnerFullName,
			OwnerPost:        e.OwnerPost,
			PassportData:     e.PassportData,
			//PictureUrl:       e.PictureUrl,
			ActivityField:    e.ActivityField,
		})
	}
	return res, nil
}

func New(conn grpc.ClientConnInterface) Client {
	return Client{grpcClient: entitiesGrpcGen.NewEntitiesClient(conn)}
}
