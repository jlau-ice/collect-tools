package repository

import (
	"context"

	"github.com/jlau-ice/collect/internal/types"
	"github.com/jlau-ice/collect/internal/types/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (s *userRepository) Create(ctx context.Context, req types.User) (id uint, err error) {
	err = s.db.WithContext(ctx).Create(&req).Error
	if err != nil {
		return 0, err
	}
	return req.Id, nil
}
