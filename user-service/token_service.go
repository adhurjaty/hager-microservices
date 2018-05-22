package main

import (
	"errors"
	pb "github.com/adhurjaty/hager-microservices/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	key = []byte("f841797abb6c3e4cf7babfb374d6f69d")
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

// Decode a token string into a token object
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	// Parse the token
	tokenType, err := jwt.ParseWithClaims(string(token), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	// Validate the token and return the custom claims
	claims, ok := tokenType.Claims.(*CustomClaims)
	if !ok || !tokenType.Valid {
		return nil, errors.New("Could not parse the claims or invalid token")
	}

	return claims, nil
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // expires tomorrow
			Issuer:    "go.micro.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
