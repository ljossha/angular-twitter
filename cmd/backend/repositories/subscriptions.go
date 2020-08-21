package repositories

import (
	"angular-twitter/cmd/backend/models"
	"context"
	"github.com/jinzhu/gorm"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *models.Subscription) (*models.Subscription, error)
	List(ctx context.Context, userID int64) ([]*models.Subscription, error)
	Delete(ctx context.Context, subscription *models.Subscription) error
}

type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (s *subscriptionRepository) Create(ctx context.Context, subscription *models.Subscription) (*models.Subscription, error) {
	err := s.db.Create(subscription).Error
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (s *subscriptionRepository) Delete(ctx context.Context, subscription *models.Subscription) error {
	return s.db.Delete(subscription).Error
}

func (s *subscriptionRepository) List(ctx context.Context, userID int64) ([]*models.Subscription, error) {
	var list []*models.Subscription
	err := s.db.Where("user_id = ?", userID).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
