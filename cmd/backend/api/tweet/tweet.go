package tweet

import (
	"angular-twitter/cmd/backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// InjectTweet inject `tweet` api group to handler
func InjectTweet(gr *gin.RouterGroup, twitterService services.TwitterService) {
	handler := gr.Group("tweet")

	handler.GET("user/:id", getTweetsByID(twitterService))
}

func getTweetsByID(twitterService services.TwitterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		tweets, err := twitterService.GetTweetsByUserID(id)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, tweets)
	}
}
