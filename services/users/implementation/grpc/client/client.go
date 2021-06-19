package client

import (
	"context"
	"github.com/ischenkx/innotech-backend/services/users"
	users2 "github.com/ischenkx/innotech-backend/services/users/implementation/grpc/pb/generated"
	"google.golang.org/grpc"
)

type Client struct {
	grpcClient users2.UsersClient
}

func (c Client) Login(ctx context.Context, dto users.LoginDto) (users.UserDto, error) {
	res, err := c.grpcClient.Login(ctx, &users2.LoginRequest{
		Username:    dto.Username,
		Password: dto.Password,
	})

	if err != nil {
		return users.UserDto{}, err
	}

	return users.UserDto{
		Username: res.Username,
		ID:       res.Id,
	}, nil
}

func (c Client) Register(ctx context.Context, dto users.RegisterDto) (users.UserDto, error) {
	res, err := c.grpcClient.Register(ctx, &users2.RegisterRequest{
		Username: dto.Username,
		Password: dto.Password,
	})

	if err != nil {
		return users.UserDto{}, err
	}

	return users.UserDto{
		Username: res.Username,
		ID:       res.Id,
	}, nil
}

func (c Client) Get(ctx context.Context, id string) (users.UserDto, error) {
	res, err := c.grpcClient.Get(ctx, &users2.GetUserRequest{Id: id})

	if err != nil {
		return users.UserDto{}, err
	}

	return users.UserDto{
		Username: res.Username,
		ID:       res.Id,
	}, nil
}

func (c Client) GetByName(ctx context.Context, name string) (users.UserDto, error) {
	res, err := c.grpcClient.GetByName(ctx, &users2.GetUserByNameRequest{Name: name})

	if err != nil {
		return users.UserDto{}, err
	}

	return users.UserDto{
		Username: res.Username,
		ID:       res.Id,
	}, nil
}

func (c Client) UpdateUsername(ctx context.Context, id, username string) error {
	_, err := c.grpcClient.UpdateUsername(ctx, &users2.UpdateUsernameRequest{
		Id:       id,
		Username: username,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c Client) UpdatePassword(ctx context.Context, id, prevPassword, password string) error {
	_, err := c.grpcClient.UpdatePassword(ctx, &users2.UpdatePasswordRequest{
		Id:               id,
		PreviousPassword: prevPassword,
		Password:         password,
	})

	if err != nil {
		return err
	}

	return nil
}
func New(client *grpc.ClientConn) Client {
	return Client{users2.NewUsersClient(client)}
}


