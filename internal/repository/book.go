package repository

import (
	"context"
	"errors"
	"time"

	"github.com/amirmahdiamini/ebook-go/internal/dto"
	"github.com/amirmahdiamini/ebook-go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	BookRepository interface {
		Insert(book *model.Book) error
		UpdateOne(updateBook dto.UpdateBook) error
		DeleteOne(sid string) error
		FindBySID(sid string) (*model.Book, error)
		GetAll() ([]dto.DBook, error)
	}
	bookRepository struct {
		database *mongo.Client
	}
)

func NewBookRepository(db *mongo.Client) BookRepository {
	return &bookRepository{
		database: db,
	}
}
func (ctx *bookRepository) Insert(book *model.Book) error {
	collection := ctx.database.Database("vbook").Collection("book")
	_, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		return errors.New("internal error #31")
	}
	return nil
}
func (ctx *bookRepository) UpdateOne(updateBook dto.UpdateBook) error {
	collection := ctx.database.Database("vbook").Collection("book")
	_, err := collection.UpdateOne(context.Background(), bson.M{"sid": updateBook.SID}, bson.M{"$set": bson.M{"title": updateBook.Title, "description": updateBook.Description}})
	if err != nil {
		return errors.New("internal error #32")
	}
	return nil
}
func (ctx *bookRepository) DeleteOne(sid string) error {
	collection := ctx.database.Database("vbook").Collection("book")
	_, err := collection.DeleteOne(context.Background(), bson.M{"sid": sid})
	if err != nil {
		return errors.New("internal error #33")
	}
	return nil
}
func (ctx *bookRepository) FindBySID(sid string) (*model.Book, error) {
	var book *model.Book
	collection := ctx.database.Database("vbook").Collection("book")
	res := collection.FindOne(context.Background(), bson.M{"sid": sid}).Decode(&book)
	if res != nil {
		return nil, errors.New("not found")
	}
	return book, nil
}
func (c *bookRepository) GetAll() ([]dto.DBook, error) {
	books := make([]dto.DBook, 0)
	collection := c.database.Database("vbook").Collection("book")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, errors.New("internal error #34")
	}
	for res.Next(ctx) {
		var book *dto.DBook
		res.Decode(&book)
		books = append(books, *book)
	}
	if err = res.Err(); err != nil {
		return nil, errors.New("دریافت اطلاعات با مشکل مواجه شد ")

	}

	return books, nil
}
