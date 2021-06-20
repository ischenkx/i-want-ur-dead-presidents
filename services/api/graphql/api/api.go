package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ischenkx/innotech-backend/services/api/graphql/graph/generated"
	"github.com/ischenkx/innotech-backend/services/api/graphql/graph/model"
	"github.com/ischenkx/innotech-backend/services/auth"
	"github.com/ischenkx/innotech-backend/services/billing"
	"github.com/ischenkx/innotech-backend/services/entities"
	dto2 "github.com/ischenkx/innotech-backend/services/entities"
	"github.com/ischenkx/innotech-backend/services/users"
	"log"
)

type API struct {
	auth     auth.Client
	users    users.Client
	entities entities.Client
	billing  billing.Client
	engine   *gin.Engine
}

func (s *API) GetEntities(ctx context.Context, req model.GetEntitiesRequest) ([]*model.Entity, error) {
	rangeDto := dto2.GetEntitiesRangeDto{
		IsPreview: *req.IsPreview,
	}
	var offset, limit int64
	if req.Offset != nil {
		offset = int64(*req.Offset)
		rangeDto.Offset = &offset
	}
	if req.Limit != nil {
		limit = int64(*req.Limit)
		rangeDto.Limit = &limit
	}
	arr, err := s.entities.GetRange(ctx, rangeDto)
	if err != nil {
		return nil, err
	}

	entis := make([]*model.Entity, len(arr))
	ids := make([]string, len(arr))
	for i, info := range arr {
		ids[i] = info.ID
		entis[i] = &model.Entity{
			Title:            info.Title,
			LongDescription:  info.LongDesc,
			ShortDescription: info.ShortDesc,
			MoneyGoal:        info.MoneyGoal,
			OwnerID:          info.OwnerID,
			ID:               info.ID,
		}
	}

	balances, err := s.billing.GetBalances(ctx, ids)
	if err != nil {
		return nil, err
	}

	for i, balance := range balances {
		entis[i].Balance = balance
	}

	return entis, nil
}

func (s *API) GetBalance(ctx context.Context, id string) (float64, error) {
	balances, err := s.billing.GetBalances(ctx, []string{id})
	if err != nil {
		return 0, err
	}
	if len(balances) < 1 {
		return 0, errors.New("failed to find balance info about the specified id")
	}
	return balances[0], nil
}

func (s *API) GetTransactions(ctx context.Context, request model.GetTransactionsRequest) ([]*model.Transaction, error) {
	user, ok := s.contextUser(ctx)

	if !ok {
		return nil, errors.New("you are not authorized")
	}

	var offset, limit *int64

	if request.Offset != nil {
		offset = new(int64)
		*offset = int64(*request.Offset)
	}

	if request.Limit != nil {
		limit = new(int64)
		*limit = int64(*request.Limit)
	}

	transactions, err := s.billing.GetTransactions(ctx, billing.GetTransactionsDto{
		Offset: offset,
		Limit:  limit,
		Id:     user.ID,
	})

	var modelTxArray []*model.Transaction

	for _, transaction := range transactions {
		modelTxArray = append(modelTxArray, &model.Transaction{
			FromID: transaction.From,
			ToID:   transaction.To,
			Amount: transaction.Amount,
		})
	}

	return modelTxArray, err
}

func (s *API) Transfer(ctx context.Context, from string, to string, amount float64) (*bool, error) {
	user, ok := s.contextUser(ctx)

	if !ok {
		return nil, errors.New("you are not authorized")
	}

	if user.ID != from {
		return nil, errors.New("you are not allowed to do request this")
	}

	return nil, s.billing.Transfer(ctx, from, to, amount)
}

func (s *API) CreateEntity(ctx context.Context, form model.EntityCreationForm) (*model.Entity, error) {
	user, ok := s.contextUser(ctx)

	if !ok {
		return nil, errors.New("you are not authorized")
	}

	shortDesc := ""

	if form.ShortDescription != nil {
		shortDesc = *form.ShortDescription
	}

	ent, err := s.entities.Create(ctx, entities.CreateEntityDto{
		Title:     form.Title,
		ShortDesc: shortDesc,
		LongDesc:  form.LongDescription,
		MoneyGoal: form.MoneyGoal,
		OwnerID:   user.ID,
	})

	if err != nil {
		return nil, err
	}

	return &model.Entity{
		Title:            ent.Title,
		LongDescription:  ent.LongDesc,
		ShortDescription: ent.ShortDesc,
		MoneyGoal:        ent.MoneyGoal,
		OwnerID:          ent.OwnerID,
		ID:               ent.ID,
		Balance:          0,
	}, nil
}

func (s *API) DeleteEntity(ctx context.Context, id string) (*model.Entity, error) {
	user, ok := s.contextUser(ctx)

	if !ok {
		return nil, errors.New("you are not authorized")
	}

	ent, err := s.entities.Delete(ctx, entities.DeleteEntityDto{
		ID:      id,
		OwnerID: &user.ID,
	})

	if err != nil {
		return nil, err
	}

	return &model.Entity{
		Title:            ent.Title,
		LongDescription:  ent.LongDesc,
		ShortDescription: ent.ShortDesc,
		MoneyGoal:        ent.MoneyGoal,
		OwnerID:          ent.OwnerID,
		ID:               ent.ID,
		Balance:          0,
	}, nil
}

func (s *API) GetUserEntities(ctx context.Context, req model.UserEntitiesRequest) ([]*model.Entity, error) {
	user, ok := s.contextUser(ctx)
	if !ok {
		return nil, errors.New("you are not authorized")
	}

	isPreview := false
	offset := (*int64)(nil)
	limit := (*int64)(nil)

	if req.IsPreview != nil {
		isPreview = *req.IsPreview
	}

	if req.Offset != nil {
		offset = new(int64)
		*offset = int64(*req.Offset)
	}

	if req.Limit != nil {
		limit = new(int64)
		*limit = int64(*req.Limit)
	}

	ents, err := s.entities.GetByOwnerID(ctx, entities.GetEntitiesByOwnerIdDto{
		OwnerID:   user.ID,
		IsPreview: isPreview,
		Offset:    offset,
		Limit:     limit,
	})

	if err != nil {
		return nil, err
	}

	var modelEnts []*model.Entity

	for _, ent := range ents {
		modelEnts = append(modelEnts, &model.Entity{
			Title:            ent.Title,
			LongDescription:  ent.LongDesc,
			ShortDescription: ent.ShortDesc,
			MoneyGoal:        ent.MoneyGoal,
			Balance:          0,
			OwnerID:          ent.OwnerID,
			ID:               ent.ID,
		})
	}

	return modelEnts, nil
}

func (s *API) UpdateEntity(ctx context.Context, updateForm model.EntityUpdateForm) (*model.Entity, error) {
	user, ok := s.contextUser(ctx)

	if !ok {
		return nil, errors.New("you are not authorized")
	}

	ents, err := s.entities.Get(ctx, entities.GetEntitiesDto{
		IDs:       []string{updateForm.ID},
		IsPreview: true,
	})

	if err != nil {
		return nil, errors.New("failed to get an entity with a specified id")
	}

	if len(ents) == 0 {
		return nil, errors.New("no entity with specified id exists")
	}

	ent := ents[0]

	if ent.OwnerID != user.ID {
		return nil, errors.New("you are not the owner")
	}

	ent, err = s.entities.Update(ctx, entities.UpdateEntityDto{
		ID:        updateForm.ID,
		Title:     updateForm.Title,
		ShortDesc: updateForm.ShortDescription,
		LongDesc:  updateForm.LongDescription,
	})

	if err != nil {
		return nil, err
	}

	return &model.Entity{
		Title:            ent.Title,
		LongDescription:  ent.LongDesc,
		ShortDescription: ent.ShortDesc,
		MoneyGoal:        ent.MoneyGoal,
		OwnerID:          ent.OwnerID,
		ID:               ent.ID,
		Balance:          0,
	}, nil
}

// utility functions

func (s *API) ginContext(ctx context.Context) (*gin.Context, bool) {
	val := ctx.Value("gin")
	ginCtx, ok := val.(*gin.Context)

	return ginCtx, ok
}

func (s *API) updateContextUser(ctx context.Context, data auth.UserData) error {
	ginCtx, ok := s.ginContext(ctx)
	if !ok {
		return errors.New("failed to get gin context")
	}
	ginCtx.Set("user", data)

	return nil
}

func (s *API) contextUser(ctx context.Context) (auth.UserData, bool) {
	ginCtx, ok := s.ginContext(ctx)
	if !ok {
		return auth.UserData{}, false
	}

	val, ok := ginCtx.Get("user")

	if !ok {
		return auth.UserData{}, false
	}

	data, ok := val.(auth.UserData)

	if !ok {
		return auth.UserData{}, false
	}

	return data, true
}

func (s *API) loadRefreshToken(ctx context.Context) (string, bool) {
	ginCtx, ok := s.ginContext(ctx)
	if !ok {
		return "", false
	}

	val, err := ginCtx.Cookie("refresh_token")

	if err != nil {
		return "", false
	}

	return val, true
}

func (s *API) loadAccessToken(ctx context.Context) (string, bool) {
	ginCtx, ok := s.ginContext(ctx)
	if !ok {
		return "", false
	}

	val, err := ginCtx.Cookie("access_token")

	if err != nil {
		return "", false
	}

	return val, true
}

func (s *API) updateRefreshToken(ctx context.Context, data string) error {
	ginCtx, ok := s.ginContext(ctx)
	if !ok {
		return errors.New("failed to update refresh token")
	}

	ginCtx.SetCookie("refresh_token", data, 0, "", "", true, true)
	return nil
}

func (s *API) updateAccessToken(ctx context.Context, data string) error {
	ginCtx, ok := s.ginContext(ctx)
	if !ok {
		return errors.New("failed to update access token")
	}

	ginCtx.SetCookie("access_token", data, 0, "", "", true, true)
	return nil
}

func (s *API) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	u, err := s.users.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Username: u.Username,
		ID:       u.ID,
	}, nil
}

func (s *API) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	u, err := s.users.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Username: u.Username,
		ID:       u.ID,
	}, nil
}

func (s *API) Login(ctx context.Context, username string, password string) (*model.User, error) {
	user, ok := s.contextUser(ctx)
	if ok {
		return &model.User{
			Username: user.Username,
			ID:       user.ID,
		}, errors.New("already authorized")
	}

	u, err := s.users.Login(ctx, users.LoginDto{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	tokens, err := s.auth.GenerateTokens(ctx, auth.UserData{
		Username: u.Username,
		ID:       u.ID,
	})

	if err != nil {
		return nil, err
	}

	if err := s.updateAccessToken(ctx, tokens.AccessToken); err != nil {
		return nil, err
	}

	if err := s.updateRefreshToken(ctx, tokens.RefreshToken); err != nil {
		return nil, err
	}

	return &model.User{
		Username: u.Username,
		ID:       u.ID,
	}, nil
}

func (s *API) Authorize(ctx context.Context) (*model.User, error) {
	u, ok := s.contextUser(ctx)

	if !ok {
		return nil, errors.New("failed to authorize: try to login")
	}

	return &model.User{
		Username: u.Username,
		ID:       u.ID,
	}, nil

}

func (s *API) Logout(ctx context.Context) (bool, error) {
	_, ok := s.contextUser(ctx)
	if err := s.updateAccessToken(ctx, ""); err != nil {
		return false, err
	}

	if err := s.updateRefreshToken(ctx, ""); err != nil {
		return false, err
	}
	return ok, nil
}

func (s *API) Register(ctx context.Context, form model.RegistrationForm) (*model.User, error) {
	user, ok := s.contextUser(ctx)
	if ok {
		return &model.User{
			Username: user.Username,
			ID:       user.ID,
		}, errors.New("already authorized")
	}

	u, err := s.users.Register(ctx, users.RegisterDto{
		Username: form.Username,
		Password: form.Password,
	})

	if err != nil {
		return nil, err
	}

	tokens, err := s.auth.GenerateTokens(ctx, auth.UserData{
		Username: u.Username,
		ID:       u.ID,
	})

	if err != nil {
		return nil, err
	}

	if err := s.updateAccessToken(ctx, tokens.AccessToken); err != nil {
		return nil, err
	}

	if err := s.updateRefreshToken(ctx, tokens.RefreshToken); err != nil {
		return nil, err
	}

	return &model.User{
		Username: u.Username,
		ID:       u.ID,
	}, nil

}

func (s *API) UpdateUsername(ctx context.Context, newUsername string) (*model.User, error) {
	user, ok := s.contextUser(ctx)

	if !ok {
		return nil, errors.New("not authorized")
	}

	if err := s.users.UpdateUsername(ctx, user.ID, newUsername); err != nil {
		return nil, err
	}

	user.Username = newUsername

	tokens, err := s.auth.GenerateTokens(ctx, user)

	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	if err := s.updateAccessToken(ctx, tokens.AccessToken); err != nil {
		return nil, err
	}

	if err := s.updateRefreshToken(ctx, tokens.RefreshToken); err != nil {
		return nil, err
	}

	if err := s.updateContextUser(ctx, user); err != nil {
		log.Println(err)
	}

	return &model.User{Username: user.Username, ID: user.ID}, nil
}

func (s *API) UpdatePassword(ctx context.Context, prevPassword string, newPassword string) (*model.User, error) {
	user, ok := s.contextUser(ctx)
	if !ok {
		return nil, errors.New("not authorized")
	}
	if err := s.users.UpdatePassword(ctx, user.ID, prevPassword, newPassword); err != nil {
		return nil, err
	}
	return &model.User{Username: user.Username, ID: user.ID}, nil
}

func (s *API) Mutation() generated.MutationResolver {
	return s
}

func (s *API) Query() generated.QueryResolver {
	return s
}

func (s *API) setup() {
	s.engine.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "gin", c))
	}, func(c *gin.Context) {

		accessToken, accessTokenRetrieved := s.loadAccessToken(c.Request.Context())

		refreshToken, refreshTokenRetrieved := s.loadRefreshToken(c.Request.Context())

		if !refreshTokenRetrieved && !accessTokenRetrieved {
			return
		}

		tokens, userData, err := s.auth.Authorize(c, auth.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := s.updateAccessToken(c.Request.Context(), tokens.AccessToken); err != nil {
			fmt.Println(err)
			return
		}

		if err := s.updateRefreshToken(c.Request.Context(), tokens.RefreshToken); err != nil {
			fmt.Println(err)
			return
		}

		if err := s.updateContextUser(c.Request.Context(), userData); err != nil {
			fmt.Println(err)
			return
		}
	})
}

func (s *API) Engine() *gin.Engine {
	return s.engine
}

func New(a auth.Client, u users.Client, e entities.Client, b billing.Client) *API {
	api := &API{auth: a, users: u, entities: e, billing: b}
	api.engine = gin.New()
	api.setup()
	return api
}
