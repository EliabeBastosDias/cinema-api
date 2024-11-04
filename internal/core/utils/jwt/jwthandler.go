package jwthandler

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/internal/envs"
)

type JWTHandler struct {
	secretKey      []byte
	userRepository ports.UserRepository
}

func New(userRepo ports.UserRepository) (*JWTHandler, error) {
	secret := envs.Get(envs.JWT_SECRET)
	if secret == "" {
		return nil, errors.New("JWT_SECRET não definido no arquivo .env")
	}

	return &JWTHandler{
		secretKey:      []byte(secret),
		userRepository: userRepo,
	}, nil
}

func (s *JWTHandler) GenerateToken(user *domain.User) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   user.Token.String(),
		ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "cinema-api",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar o token: %w", err)
	}

	return tokenString, nil
}

func (s *JWTHandler) ValidateToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao analisar o token: %w", err)
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		userToken := claims.Subject
		_, err := s.userRepository.Get(userToken)
		if err != nil {
			return nil, fmt.Errorf("usuário não encontrado: %w", err)
		}

		return claims, nil
	} else {
		return nil, errors.New("token inválido")
	}
}
