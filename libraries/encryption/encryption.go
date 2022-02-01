package encryption

import "golang.org/x/crypto/bcrypt"

type bcryptEncryption struct{}

func Bcrypt() *bcryptEncryption {
	return &bcryptEncryption{}
}

func (be *bcryptEncryption) Hash(plaintText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plaintText), 10)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (be *bcryptEncryption) Compare(hashed string, plaintext string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext)); err != nil {
		return false
	}
	return true
}
