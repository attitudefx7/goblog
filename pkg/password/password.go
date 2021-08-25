package password

import (
	"github.com/attitudefx7/goblog/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogError(err)

	return string(bytes)
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	logger.LogError(err)

	return err == nil
}

func IsHashed(str string) bool {
	// bcrypt 加密后的长度是 60
	return len(str) == 60
}
