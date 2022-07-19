package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type (
	CodeRepository interface {
		SetCode(phone string, code any, expiresIn time.Duration) error
		GetCode(phone string) (int, error)
		DeleteCode(phone string)
		IsExist(phone string) bool
		GetExpirationTime(phone string) time.Duration
	}

	codeRepository struct {
		repository *redis.Client
	}
)

func NewCodeRepository(repository *redis.Client) CodeRepository {

	return &codeRepository{
		repository: repository,
	}
}

func (c *codeRepository) SetCode(phone string, code any, expiresIn time.Duration) error {
	if err := c.repository.Set(phone, code, expiresIn).Err(); err != nil {
		fmt.Println(err)
		return errors.New("Internal Error #21")
	}
	return nil
}

func (c *codeRepository) GetCode(phone string) (int, error) {
	code, err := c.repository.Get(phone).Int()
	if err != nil {
		return 0, errors.New("Internal Error #22")
	}

	return code, nil
}

func (c *codeRepository) DeleteCode(phone string) {
	c.repository.Del(phone)

}

func (c *codeRepository) IsExist(phone string) bool {

	return c.repository.Exists(phone).Val() == 1
}

func (c *codeRepository) GetExpirationTime(phone string) time.Duration {

	return c.repository.TTL(phone).Val()
}
