package grabbing

import "context"

type UserDto struct {
	Username, Email, ID, Status string
}

type LoginDto struct {
	Username, Password string
}

type RegisterDto struct {
	Username, Email, Password string
}

type Service interface {
	Login(ctx context.Context, dto LoginDto) (UserDto, error)
	Register(ctx context.Context, dto RegisterDto) (UserDto, error)
	Get(ctx context.Context, id string) (UserDto, error)
	UpdateUsername(ctx context.Context, id, username string) error
	UpdateEmail(ctx context.Context, id, email string) error
	UpdatePassword(ctx context.Context, id, prevPassword, password string) error
	UpdatePasswordByEmail(ctx context.Context, id, password string) error
	ConfirmEmail(ctx context.Context, id string) error
	SendEmailUpdateConfirmationCode(ctx context.Context, id, code string) error
	SendPasswordUpdateConfirmation(ctx context.Context, id, code string) error
}
