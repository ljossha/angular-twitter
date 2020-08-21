package services

import (
	"angular-twitter/cmd/backend/models"
	"angular-twitter/common/config"
	"errors"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTService provides a jwt service interface
type JWTService interface {
	ParseToken(authHeader string) (*jwt.Token, error)
	GenerateToken(u *models.User) (string, error)
}

type jwtService struct {
	key  []byte
	ttl  time.Duration
	algo jwt.SigningMethod
}

// NewJWTService returns instance of jwt service
func NewJWTService() JWTService {
	return &jwtService{
		key:  []byte(config.JWTSecret()),
		algo: jwt.GetSigningMethod(config.JWTAlgorithm()),
		ttl:  time.Duration(10080) * time.Minute,
	}
}

// ParseToken parses token from Authorization header
func (s *jwtService) ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if s.algo != token.Method {
			return nil, errors.New("algorithms are different")
		}
		return s.key, nil
	})
}

// GenerateToken generates new JWT token and populates it with user data
func (s *jwtService) GenerateToken(u *models.User) (string, error) {
	return jwt.NewWithClaims(s.algo, jwt.MapClaims{
		"id":  u.ID,
		"u":   u.Name,
		"e":   u.Email,
		"exp": time.Now().Add(s.ttl).Unix(),
	}).SignedString(s.key)
}

// AuthorizedUser interface of DAL user model
type AuthorizedUser interface {
	GetUserID() int64
}

// GetUser returns the interface of authorized user
func GetUser(c *gin.Context) AuthorizedUser {
	id, _ := c.Get("id")
	name, _ := c.Get("name")
	email, _ := c.Get("email")

	return &models.User{
		ID:    int64(id.(float64)),
		Name:  name.(string),
		Email: email.(string),
	}
}
