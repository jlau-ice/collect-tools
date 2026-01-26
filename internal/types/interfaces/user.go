package interfaces

import (
	"context"

	"github.com/jlau-ice/collect/internal/types"
)

type UserService interface {
	Register(ctx context.Context, req types.User) (id uint, err error)
}

type UserRepository interface {
	Create(ctx context.Context, req types.User) (id uint, err error)
}
