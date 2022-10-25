package redigo

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var redisPool *redis.Pool

func TestMain(m *testing.M) {
	redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	defer redisPool.Close()

	retCode := m.Run()
	os.Exit(retCode)
}

func TestConn_Set(t *testing.T) {
	err := NewConn(context.Background(), redisPool.Get()).Set("k1", "v1")
	assert.NoError(t, err)
	value, err := redis.String(redisPool.Get().Do("get", "k1"))
	assert.NoError(t, err)
	assert.Equal(t, "v1", value)
}
