package services

import (
	"angular-twitter/common/config"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twitterOauth "github.com/dghubble/oauth1/twitter"
	"github.com/gin-gonic/gin"
)

// TwitterService provides a twitter service interface
type TwitterService interface {
	Login() (string, error)
	GetUser(ctx *gin.Context, requestToken string, verifier string) (*twitter.User, error)
	FindUserByName(name string) (*twitter.User, error)
	FindUserByID(id int64) (*twitter.User, error)
	GetTweetsByUserID(id int64) ([]twitter.Tweet, error)
}

type twitterService struct {
	oauth1Config *oauth1.Config
	client       *twitter.Client
}

// NewTwitterService returns instance of twitter service
func NewTwitterService() TwitterService {
	clientConfig := &oauth1.Config{
		ConsumerKey:    config.TwitterConsumerKey(),
		ConsumerSecret: config.TwitterConsumerSecret(),
		CallbackURL:    "http://localhost:4200/redirect",
		Endpoint:       twitterOauth.AuthorizeEndpoint,
	}
	token := oauth1.NewToken(config.AccessKey(), config.AccessTokenSecret())
	httpClient := clientConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return &twitterService{
		oauth1Config: clientConfig,
		client:       client,
	}
}

// Login returns link to oAuth login
func (s *twitterService) Login() (string, error) {
	requestToken, _, err := s.oauth1Config.RequestToken()
	if err != nil {
		return "", err
	}

	authorizationURL, err := s.oauth1Config.AuthorizationURL(requestToken)
	if err != nil {
		return "", err
	}

	return authorizationURL.String(), nil
}

// GetUser returns twitter from context
func (s *twitterService) GetUser(ctx *gin.Context, requestToken string, verifier string) (*twitter.User, error) {
	accessToken, accessSecret, err := s.oauth1Config.AccessToken(requestToken, s.oauth1Config.ConsumerSecret, verifier)
	if err != nil {
		return nil, err
	}

	httpClient := s.oauth1Config.Client(ctx, oauth1.NewToken(accessToken, accessSecret))
	twitterClient := twitter.NewClient(httpClient)
	accountVerifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := twitterClient.Accounts.VerifyCredentials(accountVerifyParams)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByName returns twitter user by name
func (s *twitterService) FindUserByName(name string) (*twitter.User, error) {
	user, _, err := s.client.Users.Show(&twitter.UserShowParams{
		ScreenName: name,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByID returns twitter user by id
func (s *twitterService) FindUserByID(id int64) (*twitter.User, error) {
	user, _, err := s.client.Users.Show(&twitter.UserShowParams{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetTweetsByUserID returns tweets by user id
func (s *twitterService) GetTweetsByUserID(id int64) ([]twitter.Tweet, error) {
	userTimeLine, _, err := s.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}

	return userTimeLine, nil
}
