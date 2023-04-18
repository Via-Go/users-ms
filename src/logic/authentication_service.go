package logic

import (
	"github.com/wzslr321/road_runner/server/users/mapper"
	"github.com/wzslr321/road_runner/server/users/storage"
	"golang.org/x/crypto/bcrypt"
)

type IAuthenticator interface {
	AuthenticateUser(username string, password string) *ServiceResponse
}

type Authenticator struct {
	db     storage.IUserStorage
	mapper mapper.IUserMapper
}

func NewAuthenticator(db storage.IUserStorage, mapper mapper.IUserMapper) *Authenticator {
	return &Authenticator{
		db:     db,
		mapper: mapper,
	}
}

func (s *Authenticator) AuthenticateUser(username string, password string) *ServiceResponse {
	user, err := s.db.FindUserByUsername(username)
	if err != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: err.Error(),
			Body:    []interface{}{},
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: "Bad credentials",
			Body:    []interface{}{},
		}
	}

	dto := s.mapper.MapEntityToDTO(user)

	return &ServiceResponse{
		Status:  SUCCESS,
		Message: "Logged in",
		Body:    []interface{}{dto},
	}
}
