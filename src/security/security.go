package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Verify(password string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
