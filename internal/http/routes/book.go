package routes

import (
	"github.com/amirmahdiamini/ebook-go/internal/http/book"
	"github.com/amirmahdiamini/ebook-go/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine, controller book.BookController) {
	router.MaxMultipartMemory = 9 << 20
	book := router.Group("/book")
	{
		book.Use(middleware.Authorization())
		book.POST("/add", controller.Add())
		book.GET("/all", controller.GetAll())
		book.DELETE("/delete/:sid", controller.Delete())
		book.POST("/update", controller.Update())

	}

}
