package models

// Subscription present of subscription model on DAL
type Subscription struct {
	ID             int64 `json:"id"`
	UserID         int64 `json:"user_id"`
	FollowedUserID int64 `json:"followed_user_id"`
}
