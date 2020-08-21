package main

import (
	"angular-twitter/cmd/backend/api/auth"
	"angular-twitter/cmd/backend/api/subscription"
	"angular-twitter/cmd/backend/api/tweet"
	"angular-twitter/cmd/backend/repositories"
	"angular-twitter/cmd/backend/services"
	"angular-twitter/cmd/backend/setup"
	"log"
)

// main creates and starts a Server listening.
func main() {
	handler := setup.HTTPServer()
	db := setup.DBConnect()

	gr := handler.Group("v1/api")

	userRepository := repositories.NewUserRepository(db)
	subscriptionRepository := repositories.NewSubscriptionRepository(db)

	//authService := services.NewAuthService()
	jwtService := services.NewJWTService()
	twitterService := services.NewTwitterService()
	userService := services.NewUserService(userRepository)
	subscriptionService := services.NewSubscriptionService(subscriptionRepository)

	auth.InjectAuth(gr, twitterService, jwtService, userService)
	tweet.InjectTweet(gr, twitterService)

	gr.Use(setup.AuthMiddleware(jwtService))

	subscription.InjectSubscription(gr, twitterService, subscriptionService)

	log.Fatal(handler.Run("0.0.0.0:8080"))
}
