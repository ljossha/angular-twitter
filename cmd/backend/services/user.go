package services

import (
	"angular-twitter/cmd/backend/models"
	"angular-twitter/cmd/backend/repositories"
	"context"
	"github.com/dghubble/go-twitter/twitter"
)

type UserService interface {
	Create(ctx context.Context, user *twitter.User) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Create(ctx context.Context, user *twitter.User) (*models.User, error) {
	return s.userRepository.CreateOrSelect(ctx, &models.User{
		Name:  user.Name,
		Email: user.Email,
	})
}
