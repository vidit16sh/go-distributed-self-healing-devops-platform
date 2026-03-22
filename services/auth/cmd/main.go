// entry point for the auth services
package main

import (
	"github/vidit16sh/services/auth/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", handler.HealthHandler)
	router.Run(":8081")
}
