package redigo

import (
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	Connect(Port("6380"))
	defer Close()

	conn := GetConn()
	defer conn.Close()
	do, err := conn.Do("set", "k1", "v1")
	assert.NoError(t, err)
	assert.Equal(t, "OK", do)
	value, err := redis.String(conn.Do("get", "k1"))
	assert.NoError(t, err)
	assert.Equal(t, "v1", value)
}
