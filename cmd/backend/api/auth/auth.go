package auth

import (
	"angular-twitter/cmd/backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InjectAuth inject `auth` api group to handler
func InjectAuth(gr *gin.RouterGroup, twitterService services.TwitterService, jwtService services.JWTService, userService services.UserService) {
	handler := gr.Group("auth")

	handler.GET("/login", oauthLink(twitterService))
	handler.GET("/login/redirect", redirect(twitterService, jwtService, userService))
	handler.GET("/logout", logout(twitterService))
}

type oAuthLinkResponse struct {
	URL string `json:"url"`
}

func oauthLink(twitterService services.TwitterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		url, err := twitterService.Login()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, oAuthLinkResponse{
			URL: url,
		})
	}
}

type tokenResponse struct {
	Token string `json:"token"`
}

func redirect(twitterService services.TwitterService, jwtService services.JWTService, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		twitterUser, err := twitterService.GetUser(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		user, err := userService.CreateOrSelect(c, twitterUser)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		token, err := jwtService.GenerateToken(user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, tokenResponse{
			Token: token,
		})
	}
}

func logout(twitterService services.TwitterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tweets, err := twitterService.GetTweets()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, tweets)
		return
	}
}
