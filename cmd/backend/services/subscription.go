package services

import (
	"angular-twitter/cmd/backend/models"
	"angular-twitter/cmd/backend/repositories"
	"context"
)

type SubscriptionService interface {
	AddSubscription(ctx context.Context, userID int64, subscribedUserID int64) error
	List(ctx context.Context, userID int64) ([]*models.Subscription, error)
	RemoveSubscription(ctx context.Context, userID int64, subscribedUserID int64) error
}

type subscriptionService struct {
	subscriptionRepository repositories.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepository repositories.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{
		subscriptionRepository: subscriptionRepository,
	}
}

func (s *subscriptionService) AddSubscription(ctx context.Context, userID int64, subscribedUserID int64) error {
	_, err := s.subscriptionRepository.Create(ctx, &models.Subscription{
		UserID:         userID,
		FollowedUserID: subscribedUserID,
	})

	return err
}

func (s *subscriptionService) RemoveSubscription(ctx context.Context, userID int64, subscribedUserID int64) error {
	return s.subscriptionRepository.Delete(ctx, &models.Subscription{
		UserID:         userID,
		FollowedUserID: subscribedUserID,
	})
}

func (s *subscriptionService) List(ctx context.Context, userID int64) ([]*models.Subscription, error) {
	return s.subscriptionRepository.List(ctx, userID)
}
