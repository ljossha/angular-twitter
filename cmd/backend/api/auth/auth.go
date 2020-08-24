package auth

import (
	"angular-twitter/cmd/backend/services"
	"angular-twitter/cmd/backend/setup"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InjectAuth inject `auth` api group to handler
func InjectAuth(gr *gin.RouterGroup, twitterService services.TwitterService, jwtService services.JWTService, userService services.UserService) {
	handler := gr.Group("auth")

	handler.GET("/login", oauthLink(twitterService))
	handler.POST("/login", redirect(twitterService, jwtService, userService))
	handler.GET("/me", setup.AuthMiddleware(jwtService), me())
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

type redirectRequest struct {
	Token string `json:"oauth_token"`
	Verifier string `json:"oauth_verifier"`
}

func redirect(twitterService services.TwitterService, jwtService services.JWTService, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req redirectRequest
		if err := c.Bind(&req); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		twitterUser, err := twitterService.GetUser(c, req.Token, req.Verifier)
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

func me() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
