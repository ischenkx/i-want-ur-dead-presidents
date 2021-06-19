package util

import (
	"github.com/ischenkx/innotech-backend/services/grabbing/service/db/dto"
	"github.com/ischenkx/innotech-backend/services/grabbing/service/db/models"
)

type Grabber interface {
	GrabNames(product models.Product) (dto.GrabNamesResponse, error)
	GrabFinKoefScore(product models.Product) (dto.GrabFinKoefScoreResponse, error)
	GrabCourtScore(product models.Product) (dto.GrabCourtScoreResponse, error)
	GrabSmartScore(product models.Product) (dto.GrabSmartScoreResponse, error)
}

