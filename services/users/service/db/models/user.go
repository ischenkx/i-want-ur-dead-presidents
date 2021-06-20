package models

type User struct {
	LastName  string
	FirstName string
	Email     string
	Username  string
	Password  string
	WalletID  string
	ID        string
}

type UpdateUser struct {
	LastName  *string
	FirstName *string
	Email     *string
	Username  *string
	Password  *string
	ID        string
}
