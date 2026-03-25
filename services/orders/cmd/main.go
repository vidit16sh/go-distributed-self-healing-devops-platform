package main

import (
	"github/vidit16sh/services/orders/internal/handler"

	"github.com/gin-gonic/gin"
)
func main(){ 
   router := gin.Default() 
   router.GET("/health", handler.HealthHandler)  
   router.Run(":8082")
}