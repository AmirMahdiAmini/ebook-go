package main

import (
	"net/http"
	"time"

	"github.com/amirmahdiamini/ebook-go/internal/config"
	"github.com/amirmahdiamini/ebook-go/internal/database"
	bookController "github.com/amirmahdiamini/ebook-go/internal/http/book"
	routes "github.com/amirmahdiamini/ebook-go/internal/http/routes"
	userController "github.com/amirmahdiamini/ebook-go/internal/http/user"
	"github.com/amirmahdiamini/ebook-go/internal/repository"
	bookService "github.com/amirmahdiamini/ebook-go/internal/service/book"
	userService "github.com/amirmahdiamini/ebook-go/internal/service/user"
	"github.com/amirmahdiamini/ebook-go/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ config.Config = config.New("./build/config/.env")

	db      *mongo.Client = database.SetupMongo()
	redisdb *redis.Client = database.SetupRedis()

	CodeRepository repository.CodeRepository = repository.NewCodeRepository(redisdb)

	JWT pkg.JWTService = pkg.NewJWTService()

	UserRepository repository.UserRepository     = repository.NewUserRepository(db)
	UserService    userService.UserService       = userService.NewUserService(UserRepository, CodeRepository, JWT)
	UserController userController.UserController = userController.NewUserController(UserService)

	BookRepository repository.BookRepository     = repository.NewBookRepository(db)
	BookService    bookService.BookService       = bookService.NewBookSerivce(BookRepository, UserRepository)
	BookController bookController.BookController = bookController.NewUserController(BookService)
)

func main() {
	handler := gin.New()

	routes.AuthRoutes(handler, UserController)
	routes.BookRoutes(handler, BookController)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
