package services

import (
	"angular-twitter/cmd/backend/models"
	"angular-twitter/cmd/backend/repositories"
	"context"
)

// SubscriptionService provides a subscription service interface
type SubscriptionService interface {
	AddSubscription(ctx context.Context, userID int64, subscribedUserID int64) error
	List(ctx context.Context, userID int64) ([]*SubscriptionDTO, error)
	RemoveSubscription(ctx context.Context, userID int64, subscribedUserID int64) error
}

type subscriptionService struct {
	subscriptionRepository repositories.SubscriptionRepository
}

// NewSubscriptionService returns instance of subscription service
func NewSubscriptionService(subscriptionRepository repositories.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{
		subscriptionRepository: subscriptionRepository,
	}
}

// AddSubscription 
func (s *subscriptionService) AddSubscription(ctx context.Context, userID int64, subscribedUserID int64) error {
	_, err := s.subscriptionRepository.Create(ctx, &models.Subscription{
		UserID:         userID,
		FollowedUserID: subscribedUserID,
	})

	return err
}

// RemoveSubscription perform subscription deletion
func (s *subscriptionService) RemoveSubscription(ctx context.Context, userID int64, subscribedUserID int64) error {
	return s.subscriptionRepository.Delete(ctx, &models.Subscription{
		UserID:         userID,
		FollowedUserID: subscribedUserID,
	})
}

// List returns a list of subscriptions
func (s *subscriptionService) List(ctx context.Context, userID int64) ([]*SubscriptionDTO, error) {
	subs, err := s.subscriptionRepository.List(ctx, userID)
	if err != nil {
		return nil, err
	}

	return convertToSubscriptionsDTO(subs), nil
}

// SubscriptionDTO provide subscription model on BLL
type SubscriptionDTO struct {
	UserID         int64 `json:"user_id"`
	FollowedUserID int64 `json:"followed_user_id"`
}

func convertToSubscriptionDTO(subscription *models.Subscription) *SubscriptionDTO {
	return &SubscriptionDTO{
		UserID: subscription.UserID,
		FollowedUserID: subscription.FollowedUserID,
	}
}

func convertToSubscriptionsDTO(subscriptions []*models.Subscription) []*SubscriptionDTO {
	var res []*SubscriptionDTO
	for _, val := range subscriptions {
		res = append(res, convertToSubscriptionDTO(val))
	}

	return res
}
