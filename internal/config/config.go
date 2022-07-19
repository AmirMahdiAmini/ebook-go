package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type config struct{}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}

func New(path string) Config {
	Parse(path)
	return &config{}
}
func configENV(fileNames ...string) error {
	err := godotenv.Load(fileNames...)
	if err != nil {
		return err
	}
	return nil
}
func Parse(path string) {
	a := fileExtension(path)
	switch a {
	// case "yaml":
	case "env":
		configENV(path)
	}
}

func fileExtension(path string) string {
	split := strings.Split(path, ".")
	return split[len(split)-1]
}
