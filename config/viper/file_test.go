package viper

import (
	"github.com/hlf513/go-pkg/config"
	"github.com/magiconair/properties/assert"
	"testing"
)

var user User

type User struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

type Config struct {
	User User `json:"user"`
}

func (c Config) Init() error {
	user = c.User
	return nil
}

func (c Config) Close() error {
	return nil
}

func TestNewFile(t *testing.T) {
	conf := NewFile(
		config.Name("file_test"),
		config.Path("."),
		config.AddConfig(
			new(Config),
		),
		config.Watcher(false),
	)
	defer conf.Close()

	assert.Equal(t, user.Name, "john")
	assert.Equal(t, user.Age, 1)
	assert.Equal(t, user.Height, 10.23)
}
