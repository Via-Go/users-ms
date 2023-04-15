package domain

import (
	"github.com/golang-jwt/jwt"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type JwtClaims struct {
	User *pb.UserDTO
	jwt.StandardClaims
}

func NewJwtClaims(user *pb.UserDTO, claims jwt.StandardClaims) *JwtClaims {
	return &JwtClaims{
		User:           user,
		StandardClaims: claims,
	}
}
