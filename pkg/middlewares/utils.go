package middlewares

import (
	"fmt"
	"math/rand"
	"strings"

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

func GenerateSubBatchCodes(batchCode string, numberOfSubBatches int) []string {
	var subBatchCodes []string
	for i := 1; i <= numberOfSubBatches; i++ {
		subBatchCodes = append(subBatchCodes, fmt.Sprintf("%s/%s", batchCode, strings.ToUpper(getAlphabetRepresentation(i))))
		continue
	}
	return subBatchCodes
}

func getAlphabetRepresentation(n int) string {
	var result strings.Builder

	for n > 0 {
		remainder := n % 26

		if remainder == 0 {
			remainder = 26
			n--
		}

		char := 'a' + rune(remainder-1)

		result.WriteByte(byte(char))

		n /= 26
	}

	reversed := result.String()
	runes := []rune(reversed)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
