package encryption

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct {}

func (h BcryptHasher) Verify(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func (h BcryptHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

