package controller

import (
	"go-login-crud/dto"
	"go-login-crud/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func CreateAuthController(s service.AuthService) *AuthController {
	return &AuthController{s}
}

func (u *AuthController) Login(c *gin.Context) {
	var req dto.AuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result, err := u.authService.Auth(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, result)
}