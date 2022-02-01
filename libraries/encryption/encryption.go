package encryption

import "golang.org/x/crypto/bcrypt"

type BcryptEncryption struct{}

func Bcrypt() *BcryptEncryption {
	return &BcryptEncryption{}
}

func (be *BcryptEncryption) Hash(plaintText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plaintText), 10)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (be *BcryptEncryption) Compare(hashed string, plaintext string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext)); err != nil {
		return false
	}
	return true
}
