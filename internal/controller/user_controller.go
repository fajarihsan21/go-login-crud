package controller

import (
	"go-login-crud/dto"
	"go-login-crud/internal/service"
	"go-login-crud/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserController struct {
	userService service.UserService
}

func CreateUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (u *UserController) UserRegister(c *gin.Context) {
	var req dto.UserRegist
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := u.userService.Register(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := dto.RegistRes{
		Message: "Success",
		Code: "200",
		Result: respone,
	}

	c.JSON(200, res)
}

func (u *UserController) UserUpdate(c *gin.Context) {
	var req dto.UserData
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("error: %s", err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := u.userService.Update(req)
	if err != nil {
		log.Printf("error: %s", err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := dto.RegistRes{
		Message: "Success",
		Code: "200",
		Result: respone,
	}

	c.JSON(200, res)
}

func (u *UserController) UserDelete(c *gin.Context) {
	var req dto.DelReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := u.userService.Delete(req.Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := dto.RegistRes{
		Message: "Success",
		Code: "200",
		Result: "User Deleted",
	}

	c.JSON(200, res)
}

func (u *UserController) GetAllData(c *gin.Context) {
	var req dto.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := u.userService.FetchData(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := model.UsersRes{
		Message: "Success",
		Code: "200",
		Result: respone,
	}

	c.JSON(200, res)
}