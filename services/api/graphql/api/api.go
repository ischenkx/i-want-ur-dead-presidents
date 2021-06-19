package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ischenkx/innotech-backend/services/api/graphql/graph/generated"
	"github.com/ischenkx/innotech-backend/services/api/graphql/graph/model"
	"github.com/ischenkx/innotech-backend/services/auth"
	"github.com/ischenkx/innotech-backend/services/users"
	"log"
)

type API struct {
	auth auth.Client
	users users.Client
	engine *gin.Engine
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

func New(a auth.Client, u users.Client) *API {
	api := &API{auth:  a, users: u}
	api.engine = gin.New()
	api.setup()
	return api
}