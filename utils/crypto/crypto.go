package crypto

import (
	"context"
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const HashingCost int = 4

type cryptoService struct {
	bcryptHashingCost int
}

func NewCryptoService() CryptoService {
	return &cryptoService{
		bcryptHashingCost: HashingCost,
	}
}

// CreatePasswordHash Create a password hash of given plain password.
func (s *cryptoService) CreatePasswordHash(ctx context.Context, plainPassword string) (hashedPassword string, err error) {
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), s.bcryptHashingCost)
	if err != nil {
		return
	}

	hashedPassword = string(passwordHashInBytes)
	return
}

// ValidatePassword Validate given plain password against hashed password. It return true if matched.
func (s *cryptoService) ValidatePassword(ctx context.Context, hashedPassword, plainPassword string) (isValid bool) {
	hashedPasswordInBytes := []byte(hashedPassword)
	plainPasswordInBytes := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordInBytes, plainPasswordInBytes)
	isValid = err == nil
	return
}

// CreateMD5Hash return md5 hash value of given plain text
func (s *cryptoService) CreatedMD5Hash(ctx context.Context, plainText string) (hashedText string) {
	plainTextInBytes := []byte(plainText)
	resultInByte := md5.Sum(plainTextInBytes)
	hashedText = fmt.Sprintf("%x", resultInByte)
	return
}
