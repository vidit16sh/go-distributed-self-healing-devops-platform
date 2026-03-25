// entry point for the auth services
package main

import (
	"github/vidit16sh/services/auth/internal/handler"
	"github/vidit16sh/services/auth/internal/middleware"
	"github/vidit16sh/services/auth/internal/repository"
	"github/vidit16sh/services/auth/internal/service" 

	"github.com/gin-gonic/gin"
)

func main(){ 
    router := gin.Default() 
	
	userRepo := repository.NewUserRepository() // creating an instance of repo usingg constructor
	authService := service.NewAuthService(userRepo) // similarly for service 
	authHandler := handler.NewAuthHandler(authService) // and for handler also 

	authRouter := router.Group("/auth") 
	{ 
		authRouter.POST("/register",authHandler.Register) 
		authRouter.POST("/login",authHandler.Login)
	} 
	
	protected := router.Group("/auth")  // group to apply middlware to set of routes 
    protected.Use(middleware.JWTMiddleware())
	{
		protected.GET("/profile", authHandler.Profile)
	}

	router.GET("/health", handler.HealthHandler) 
	
	router.Run(":5000")
}
