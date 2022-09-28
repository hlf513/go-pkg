package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect(
		//Host("127.0.0.1"),
		//Username("user"),
		//Password("pwd"),
	)
	if err != nil {
		assert.Error(t, err)
		return
	}
	var result int
	db.Select("1+1").Table("user").Find(&result)
	assert.Equal(t, 2, result)
}
