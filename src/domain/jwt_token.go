package domain

import (
	pb "buf.build/gen/go/viago/users-ms/protocolbuffers/go/v1"
	"github.com/golang-jwt/jwt"
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
