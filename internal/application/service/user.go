package service

import (
	"context"

	"github.com/jlau-ice/collect/internal/types"
	"github.com/jlau-ice/collect/internal/types/interfaces"
)

type userService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(
	userRepository interfaces.UserRepository,
) interfaces.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Register(ctx context.Context, req types.User) (id uint, err error) {
	return s.userRepository.Create(ctx, req)
}
