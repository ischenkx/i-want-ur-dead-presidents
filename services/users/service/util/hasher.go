package util

type PasswordHasher interface {
	Verify(hash string, src string) error
	Hash(string) (string, error)
}

