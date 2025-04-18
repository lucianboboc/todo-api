package jsonwebtoken

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenString = string

type Service interface {
	GenerateToken(userID int, exp time.Duration) (TokenString, error)
	// ValidateToken returns user_id if token is valid
	ValidateToken(token string) (int, error)
}

func NewService(jwtSecret string) Service {
	return service{
		secret: jwtSecret,
	}
}

type service struct {
	secret string
}

func (s service) GenerateToken(userID int, exp time.Duration) (TokenString, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(exp).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s service) ValidateToken(token string) (int, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return 0, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	userID, ok := claims["sub"].(int)
	if !ok {
		return 0, errors.New("Invalid claims sub entry")
	}

	return userID, nil
}
