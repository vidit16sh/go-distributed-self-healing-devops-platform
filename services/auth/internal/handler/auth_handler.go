package handler 

import (
	"net/http"

	"github/vidit16sh/services/auth/internal/service"


	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}


func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
} 

func HealthHandler(c *gin.Context) { // c is standard variable and *gin.Context refer to the request/response context 
	c.JSON(http.StatusOK, gin.H{
		"status": "auth service running",
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	} 
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	err := h.service.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} 
	c.JSON(http.StatusOK, gin.H{"message": "registration successful"}) 
}