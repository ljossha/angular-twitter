package subscription

import (
	"angular-twitter/cmd/backend/services"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// InjectSubscription inject `subscribe` api group to handler
func InjectSubscription(gr *gin.RouterGroup, twitterService services.TwitterService, subscriptionService services.SubscriptionService) {
	handler := gr.Group("subscription")

	handler.POST("", subscribe(twitterService, subscriptionService))
	handler.DELETE("/:id", unSubscribe(twitterService, subscriptionService))
	handler.GET("", list(twitterService, subscriptionService))
}

type subscribeRequest struct {
	Name string `json:"name"`
}

func subscribe(twitterService services.TwitterService, subscriptionService services.SubscriptionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req subscribeRequest
		if err := c.Bind(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		authorizedUser := services.GetUser(c)

		twitterUser, err := twitterService.FindUserByName(req.Name)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err = subscriptionService.AddSubscription(c, authorizedUser.ID, twitterUser.ID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func list(twitterService services.TwitterService, subscriptionService services.SubscriptionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizedUser := services.GetUser(c)

		subscriptions, err := subscriptionService.List(c, authorizedUser.ID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var listUsers []*twitter.User

		for _, sub := range subscriptions {
			user, err := twitterService.FindUserByID(sub.FollowedUserID)
			if err == nil {
				listUsers = append(listUsers, user)
			}
		}

		c.JSON(http.StatusOK, listUsers)
	}
}

func unSubscribe(twitterService services.TwitterService, subscriptionService services.SubscriptionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		authorizedUser := services.GetUser(c)
		twitterUser, err := twitterService.FindUserByID(id)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err = subscriptionService.RemoveSubscription(c, authorizedUser.ID, twitterUser.ID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Status(http.StatusOK)
	}
}
