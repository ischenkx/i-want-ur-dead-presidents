package validation

import "net/mail"

type Validator struct {}

func (v Validator) ValidateUsername(s string) error {
	return nil
}

func (v Validator) ValidateEmail(s string) error {
	_, err := mail.ParseAddress(s)
	return err
}

func (v Validator) ValidatePassword(s string) error {
	return nil
}

