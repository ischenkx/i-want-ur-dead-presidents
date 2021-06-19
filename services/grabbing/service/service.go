package service

import (
	"context"
	db "github.com/ischenkx/innotech-backend/services/grabbing/service/db"
	models "github.com/ischenkx/innotech-backend/services/grabbing/service/db/models"
	util "github.com/ischenkx/innotech-backend/services/grabbing/service/util"
)

type Service struct {
	db      db.DB
	grabber util.Grabber
}

func (s *Service) Get(ctx context.Context, p models.Product) (models.Response, error) {
	names, err := s.grabber.GrabNames(p)
	if err != nil {
		return models.Response{}, err
	}

	courtScore, err := s.grabber.GrabCourtScore(p)
	if err != nil {
		return models.Response{}, err
	}

	smartScore, err := s.grabber.GrabSmartScore(p)
	if err != nil {
		return models.Response{}, err
	}

	finKoefScore, err := s.grabber.GrabFinKoefScore(p)
	if err != nil {
		return models.Response{}, err
	}

	return models.Response{
		Id:       p.Id,
		Inn:      p.Inn,
		Name:     names.Name,
		FullName: names.FullName,
		Score: models.Score{
			OverallScore: courtScore.Score + smartScore.Score + finKoefScore.Score,
			CourtScore:   courtScore.Score,
			FinKoefScore: smartScore.Score,
			SmartScore:   finKoefScore.Score,
		},
	}, nil
}


func New(db db.DB, grabber util.Grabber) *Service {
	return &Service{
		db:        db,
		grabber: grabber,
	}
}
