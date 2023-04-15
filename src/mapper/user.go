package mapper

import (
	"github.com/wzslr321/road_runner/server/users/src/domain"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type IUserMapper interface {
	MapEntityToDTO(entity *domain.User) *pb.UserDTO
}

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) MapEntityToDTO(entity *domain.User) *pb.UserDTO {
	return &pb.UserDTO{
		Id:       entity.Id,
		Username: entity.Username,
		Email:    entity.Email,
	}
}
