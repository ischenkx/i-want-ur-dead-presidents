package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ischenkx/innotech-backend/services/users/service/db"
	"github.com/ischenkx/innotech-backend/services/users/service/db/models"
	"github.com/ischenkx/innotech-backend/services/users/service/util"
)

type Service struct {
	db                       db.DB
	hasher                   util.PasswordHasher
	validator                util.Validator
}

func (s *Service) Login(ctx context.Context, username, password string) (models.User, error) {
	user, err := s.db.GetUserByName(ctx, username)
	if err != nil {
		return models.User{}, err
	}
	if err := s.hasher.Verify(user.Password, password); err != nil {
		return models.User{}, err
	}
	return user, err
}

func (s *Service) Register(ctx context.Context, username, password string) (models.User, error) {
	hashedPassword, err := s.hasher.Hash(password)

	if err != nil {
		return models.User{}, err
	}

	user, err := s.db.CreateUser(ctx, models.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: hashedPassword,
	})

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *Service) UpdateUsername(ctx context.Context, id string, username string) error {
	if s.validator != nil {
		err := s.validator.ValidateUsername(username)
		if err != nil {
			return err
		}
	}
	_, err := s.db.UpdateUser(ctx, models.UpdateUser{
		Username: &username,
		ID:       id,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdatePassword(ctx context.Context, id string, previousPassword string, password string) error {
	if s.validator != nil {
		err := s.validator.ValidatePassword(password)
		if err != nil {
			return err
		}
	}

	u, err := s.db.GetUser(ctx, id)

	if err != nil {
		return err
	}

	if err := s.hasher.Verify(u.Password, previousPassword); err != nil {
		return fmt.Errorf("failed to verify previous password: %s", previousPassword)
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}

	_, err = s.db.UpdateUser(ctx, models.UpdateUser{
		Password: &hash,
		ID:       id,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetByName(ctx context.Context, username string) (models.User, error) {
	return s.db.GetUserByName(ctx, username)
}

func (s *Service) Get(ctx context.Context, id string) (models.User, error) {
	return s.db.GetUser(ctx, id)
}

func New(db db.DB, validator util.Validator, hasher util.PasswordHasher) *Service {
	return &Service{
		db:                       db,
		hasher:                   hasher,
		validator:                validator,
	}
}
