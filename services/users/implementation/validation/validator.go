package validation

import (
	"errors"
	"net/mail"
)

type Validator struct {}

func (v Validator) ValidateUsername(s string) error {
	if len(s) < 4 {
		return errors.New("username should be longer than 4 symbols")
	}

	if len(s) > 16 {
		return errors.New("username should be less than 16 symbols")
	}

	return nil
}

func (v Validator) ValidateEmail(s string) error {
	_, err := mail.ParseAddress(s)
	return err
}

func (v Validator) ValidatePassword(s string) error {
	if len(s) < 4 {
		return errors.New("password should be longer than 4 symbols")
	}

	if len(s) > 16 {
		return errors.New("password should be less than 16 symbols")
	}

	return nil
}

