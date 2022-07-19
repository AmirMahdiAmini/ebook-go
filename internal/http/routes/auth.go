package routes

import (
	"net/http"

	"github.com/amirmahdiamini/ebook-go/internal/http/user"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, controller user.UserController) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controller.Signup())
		auth.POST("/signin", controller.Signin())
		auth.POST("/verify", controller.VerifyAccount())
		auth.POST("/forgot_password", controller.ForgotPassword())
		auth.POST("/change_password", controller.ChangePassword())
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"response": "PONG",
			})
		})
	}

}
