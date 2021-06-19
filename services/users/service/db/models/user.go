package models

type User struct {
	Username string
	Password string
	ID       string
}

type UpdateUser struct {
	Username *string
	Password *string
	ID       string
}
