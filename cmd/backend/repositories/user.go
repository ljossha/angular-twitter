package repositories

import (
	"angular-twitter/cmd/backend/models"
	"context"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateOrSelect(ctx context.Context, user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (s *userRepository) CreateOrSelect(ctx context.Context, user *models.User) (*models.User, error) {
	err := s.db.Where("email = ?", user.Email).First(user).Error
	if err == nil {
		return user, nil
	}

	err = s.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
