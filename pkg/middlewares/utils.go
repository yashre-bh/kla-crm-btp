package middlewares

import (
	"fmt"
	"math/rand"

	"github.com/yashre-bh/kla-crm-btp/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special = "!@#$%&*()_+=[]{}<>?~"
	numbers = "0123456789"
)

func GenerateRandomPassword(length int, useLetters bool, useSpecial bool, useNum bool) string {
	b := make([]byte, length)
	for i := range b {
		if useLetters {
			b[i] = letters[rand.Intn(len(letters))]
		} else if useSpecial {
			b[i] = special[rand.Intn(len(special))]
		} else if useNum {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
	}
	return string(b)
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func CompareHashedPasswords(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func CreateBatchCode(entity string, date string) (string, error) {
	entityCode, err := models.GetEntityCode(entity)
	if err != nil || len(entityCode) == 0 {
		return "", err
	}

	batchCode := fmt.Sprintf("%s/%s", entityCode, date)

	return batchCode, nil
}
