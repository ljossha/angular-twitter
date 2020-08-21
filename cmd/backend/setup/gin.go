package setup

import (
	"angular-twitter/cmd/backend/services"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HTTPServer returns a configured gin engine
func HTTPServer() *gin.Engine {
	handler := gin.New()

	serveWebApp(handler)

	handler.Use(gin.Recovery())

	handler.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	})

	return handler
}

// We should serve SPA frontend
func serveWebApp(handler *gin.Engine) {
	handler.NoRoute(func(context *gin.Context) {
		context.File("/frontend/build/index.html")
	})
}

// AuthMiddleware require authorization for endpoint
func AuthMiddleware(service services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.GetHeader("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		response, err := service.ParseToken(reqToken)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		if !response.Valid {
			c.Status(http.StatusForbidden)
			c.Abort()
		}

		claims := response.Claims.(jwt.MapClaims)

		id := claims["id"]
		name := claims["u"]
		email := claims["e"]

		c.Set("id", id)
		c.Set("name", name)
		c.Set("email", email)
		c.Next()
	}
}
