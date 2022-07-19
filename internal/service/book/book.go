package book

import (
	"errors"

	"github.com/amirmahdiamini/ebook-go/internal/dto"
	"github.com/amirmahdiamini/ebook-go/internal/model"
	"github.com/amirmahdiamini/ebook-go/internal/repository"
)

type (
	BookService interface {
		Add(book *model.Book) error
		UpdateOne(updateBook dto.UpdateBook, user_sid string) error
		Delete(sid string, user_sid string) error
		GetAll() ([]dto.DBook, error)
	}
	bookService struct {
		bookRepository repository.BookRepository
		userRepository repository.UserRepository
	}
)

func NewBookSerivce(bookRepository repository.BookRepository, userRepository repository.UserRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
		userRepository: userRepository,
	}
}

func (ctx *bookService) Add(book *model.Book) error {
	err := ctx.bookRepository.Insert(book)
	if err != nil {
		return err
	}
	return nil
}
func (ctx *bookService) UpdateOne(updateBook dto.UpdateBook, user_sid string) error {
	user, err := ctx.userRepository.FindBySID(user_sid)
	if err != nil {
		return err
	}
	book, err := ctx.bookRepository.FindBySID(updateBook.SID)
	if err != nil {
		return err
	}
	if user.SID != book.UserSID {
		return errors.New("your are not owner of this book")
	}

	err = ctx.bookRepository.UpdateOne(updateBook)
	if err != nil {
		return err
	}
	return nil
}
func (ctx *bookService) Delete(sid string, user_sid string) error {
	user, err := ctx.userRepository.FindBySID(user_sid)
	if err != nil {
		return err
	}
	book, err := ctx.bookRepository.FindBySID(sid)
	if err != nil {
		return err
	}
	if user.SID != book.UserSID {
		return errors.New("your are not owner of this book")
	}

	err = ctx.bookRepository.DeleteOne(sid)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *bookService) GetAll() ([]dto.DBook, error) {
	books, err := ctx.bookRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}
