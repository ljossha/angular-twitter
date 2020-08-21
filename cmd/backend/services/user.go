package services

import (
	"angular-twitter/cmd/backend/models"
	"angular-twitter/cmd/backend/repositories"
	"context"
	"github.com/dghubble/go-twitter/twitter"
)

// UserService provides a user service interface
type UserService interface {
	CreateOrSelect(ctx context.Context, user *twitter.User) (*UserDTO, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

// NewUserService returns instance of user service
func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// CreateOrSelect returns created user or just return already created user
func (s *userService) CreateOrSelect(ctx context.Context, user *twitter.User) (*UserDTO, error) {
	createdUser, err := s.userRepository.CreateOrSelect(ctx, &models.User{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}

	return convertToUserDTO(createdUser), nil
}

// UserDTO provide user model on BLL
type UserDTO struct {
	ID    int64  `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func convertToUserDTO(user *models.User) *UserDTO {
	return &UserDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
}
