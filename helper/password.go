package helper

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	// IterationsCount to set the iteration
	IterationsCount = 15000
	// SaltSize to set salt size
	SaltSize = 64
	// KeyLength to set key length
	KeyLength = 64
	// PasswordCharSet to set password charset
	PasswordCharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789/*-+!@#$%^&*()-_~`|"
	// RandomPasswordLength to set length of random password
	RandomPasswordLength = 10
)

// GenerateSalt function for generating salt
func GenerateSalt() (saltBytes []byte, saltString string) {
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err == nil {
		saltBytes = salt
		saltString = strconv.Itoa(IterationsCount) + "." + base64.StdEncoding.EncodeToString(salt)
	} else {
		saltString = ""
	}
	return
}

// CalculatePasswordHash function for calculating result password hash
// password string plain password that will be stored to DB
// salt string salt for generating password
func CalculatePasswordHash(password string, salt string) string {
	if salt == "" {
		saltBytes, _ := GenerateSalt()
		pass := base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(password), saltBytes, IterationsCount, KeyLength, sha1.New))
		return pass
	}

	salts := strings.Split(salt, ".")
	iterations, _ := strconv.Atoi(salts[0])
	saltBytes, _ := base64.StdEncoding.DecodeString(salts[1])
	pass := base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(password), saltBytes, iterations, KeyLength, sha1.New))

	return pass
}

// CheckPassword function for checking password
// plain string plain password
// password string generated and saved password from DB
// salt string salt for generating password
func CheckPassword(plain, password, salt string) bool {
	// replace space while checking because inserting process causes so much spaces
	generatedPass := strings.Replace(CalculatePasswordHash(plain, salt), " ", "", -1)
	fromDBPass := strings.Replace(password, " ", "", -1)

	if generatedPass == fromDBPass {
		return true
	}
	return false
}

// GenerateRandomPassword function for generating random password
// length int random passwor length
func GenerateRandomPassword(length int) (plaintext string, salt string, ciphertext string) {
	newPassword := make([]byte, length)
	_, err := rand.Read(newPassword)
	if err != nil {
		return "", "", ""
	}
	// iterate new password
	for i := range newPassword {
		newPassword[i] = PasswordCharSet[int(newPassword[i])%len(PasswordCharSet)]
	}
	plaintext = string(newPassword)
	_, salt = GenerateSalt()
	ciphertext = CalculatePasswordHash(plaintext, salt)

	return
}
