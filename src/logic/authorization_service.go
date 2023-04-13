package logic

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/wzslr321/road_runner/server/users/src/domain"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"strings"
	"time"
)

type IAuthorizer interface {
	GenerateTokens(userDTO *pb.UserDTO) (map[string]string, error)
	DecodeToken(raw string) (*domain.JwtClaims, error)
	isTokenValid(raw string) bool
	RefreshTokens(rawRefresh string) (map[string]string, error)
}

type Authorizer struct {
}

func NewAuthorizer() *Authorizer {
	return &Authorizer{}
}

func (s *Authorizer) GenerateTokens(userDTO *pb.UserDTO) (map[string]string, error) {
	baseTokenClaims := domain.NewJwtClaims(
		userDTO,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		})

	refreshTokenClaims := domain.NewJwtClaims(
		userDTO,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	)

	rawBaseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, baseTokenClaims)
	signedBaseToken, err := rawBaseToken.SignedString([]byte("top-secret-secret"))
	if err != nil {
		return nil, err
	}

	rawRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, err := rawRefreshToken.SignedString([]byte("top-secret-secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"Base token":    "Bearer " + signedBaseToken,
		"Refresh token": "Bearer " + signedRefreshToken,
	}, nil
}

func (s *Authorizer) DecodeToken(raw string) (*domain.JwtClaims, error) {
	if !s.isTokenValid(raw) {
		return nil, fmt.Errorf("invalid token")
	}

	token, err := jwt.ParseWithClaims(raw, &domain.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token parsing failed")
	}

	return token.Claims.(*domain.JwtClaims), nil
}

func (s *Authorizer) isTokenValid(raw string) bool {
	if strings.TrimSpace(raw) == "" {
		return false
	}

	if !strings.HasPrefix(raw, "Bearer ") {
		return false
	}

	return true
}

func (s *Authorizer) RefreshTokens(rawRefresh string) (map[string]string, error) {
	decodedToken, err := s.DecodeToken(rawRefresh)
	if err != nil {
		return nil, err
	}

	tokens, err := s.GenerateTokens(decodedToken.User)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
