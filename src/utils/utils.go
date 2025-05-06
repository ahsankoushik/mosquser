package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key-yoooo")

type UserClaims struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Logger(msg ...any) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("[%s:%d]\n", file, line)
	log.Println(msg)
}

func FormatValidationError(err error) (string, map[string]string) {
	data := make(map[string]string)
	msg := "Validation error on "
	for _, err := range err.(validator.ValidationErrors) {
		msg += err.Field() + " " + err.Tag() + ", "
		data[strings.ToLower(err.Field())] = err.Tag()
	}

	return msg, data
}

func CreateTokenWithTime(id uint, username string, days int) (string, error) {
	// Set custom claims
	claims := UserClaims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(days))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token with the specified claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateToken(id uint, username string) (string, error) {
	return CreateTokenWithTime(id, username, 7)
}

func VerifyToken(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		b[i] = letterRunes[randIndex.Int64()]
	}
	return string(b)
}

func RandomInt(n int) string {
	otpStr := ""
	max := big.NewInt(9)
	for i := 0; i < n; i++ {
		randNum, _ := rand.Int(rand.Reader, max)
		otpStr += strconv.FormatInt(randNum.Int64(), 10)
	}
	return otpStr
}
