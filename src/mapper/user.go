package mapper

import (
	pb "buf.build/gen/go/viago/auth/protocolbuffers/go"
	"github.com/wzslr321/road_runner/server/users/domain"
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
