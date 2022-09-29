package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	err := Connect()
	assert.NoError(t, err)

	var result int
	GetDB().Select("1+1").Table("user").Find(&result)
	assert.Equal(t, 2, result)
}
