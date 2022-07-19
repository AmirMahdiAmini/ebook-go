package user

import (
	"fmt"
	"net/http"

	"github.com/amirmahdiamini/ebook-go/internal/dto"
	"github.com/amirmahdiamini/ebook-go/internal/model"
	"github.com/amirmahdiamini/ebook-go/internal/service/user"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Signup() gin.HandlerFunc
		Signin() gin.HandlerFunc
		VerifyAccount() gin.HandlerFunc
		ForgotPassword() gin.HandlerFunc
		ChangePassword() gin.HandlerFunc
	}
	userController struct {
		service user.UserService
	}
)

func NewUserController(service user.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (sv *userController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "send the correct data",
			})
			return
		}
		if err := sv.service.Signup(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "user created",
		})
	}
}

func (sv *userController) Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data dto.Signin
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "send the correct data",
			})
			return
		}
		token, err := sv.service.Signin(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
func (sv *userController) VerifyAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data dto.VerifyAccount
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "send the correct data",
			})
			return
		}
		if err := sv.service.VerifyAccount(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
func (sv *userController) ForgotPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data dto.ForgotPassword
		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "send the correct data",
			})
			return
		}
		if err := sv.service.ForgotPassword(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "password code sent your phone",
		})
	}
}

func (sv *userController) ChangePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"Post": "ChangePassword",
		})
	}
}
