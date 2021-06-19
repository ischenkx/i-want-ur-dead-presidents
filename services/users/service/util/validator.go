package util

type Validator interface {
	ValidateUsername(string) error
	ValidatePassword(string) error
}
