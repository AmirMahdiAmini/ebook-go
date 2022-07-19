package book

import (
	"net/http"

	"github.com/amirmahdiamini/ebook-go/internal/dto"
	"github.com/amirmahdiamini/ebook-go/internal/model"
	"github.com/amirmahdiamini/ebook-go/internal/service/book"
	"github.com/amirmahdiamini/ebook-go/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	BookController interface {
		Add() gin.HandlerFunc
		Delete() gin.HandlerFunc
		GetAll() gin.HandlerFunc
		Update() gin.HandlerFunc
	}
	bookController struct {
		service book.BookService
	}
)

func NewUserController(service book.BookService) BookController {
	return &bookController{
		service: service,
	}
}
func (sv *bookController) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book
		book.Title = c.PostForm("title")
		if book.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "title is required",
			})
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "send the correct data",
			})
			return
		}
		if file.Header.Get("Content-Type") != "application/pdf" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "send the correct data",
			})
			return
		}
		path := "./uploads/" + uuid.New().String()
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error in saving file",
			})
			return
		}

		book.Path = path
		book.SID = pkg.SID(18)
		book.Description = c.PostForm("description")
		book.UserSID = c.MustGet("sid").(string)
		if err := sv.service.Add(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "book created",
		})
	}
}
func (sv *bookController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Param("id")
		user_sid := c.MustGet("sid").(string)
		if err := sv.service.Delete(sid, user_sid); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "book deleted",
		})
	}
}
func (sv *bookController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := sv.service.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	}
}

func (sv *bookController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book dto.UpdateBook
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user_sid := c.MustGet("sid").(string)
		if err := sv.service.UpdateOne(book, user_sid); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "book updated",
		})

	}
}
