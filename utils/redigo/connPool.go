package redigo

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
)

type Conn interface {
	Close() // 连接放入连接池
	Set(key string, value interface{}) error
	SetExpire(key string, value interface{}, expire int) error
	GetString(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	Del(key string) error
	Lock(key string, timeout int) (bool, error)
	Unlock(key string) error
	Incr(key string) (int64, error)
	Expire(key string, expire int) error
	LPush(key string, value interface{}) error
	LPop(key string) (interface{}, error)
	RPush(key string, value interface{}) error
	RPop(key string) (string, error)
	Do(commandName string, key string, args ...interface{}) (interface{}, error)
}

func NewConn(ctx context.Context, c redis.Conn) Conn {
	var span opentracing.Span
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		span, _ = opentracing.StartSpanFromContext(ctx, "Redis")
	}

	return conn{
		conn: c,
		span: span,
	}
}

type conn struct {
	conn redis.Conn
	span opentracing.Span
}

func (r conn) Close() {
	if r.span != nil {
		r.span.Finish()
	}
	_ = r.conn.Close()
}

func (r conn) Set(key string, value interface{}) error {
	if r.span != nil {
		r.span.LogKV("command", "set", "key", key, "value", value)
	}
	if _, err := r.conn.Do("Set", key, value); err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return err
	}
	return nil
}

func (r conn) SetExpire(key string, value interface{}, expire int) error {
	if r.span != nil {
		r.span.LogKV("command", "setExpire", "key", key, "value", value, "expire", expire)
	}
	if _, err := r.conn.Do("Set", key, value, "EX", expire); err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return err
	}
	return nil
}

func (r conn) GetString(key string) (string, error) {
	if r.span != nil {
		r.span.LogKV("command", "getString", "key", key)
	}
	rep, err := redis.String(r.conn.Do("Get", key))
	if err == redis.ErrNil {
		return "", nil
	}
	if err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return "", err
	}
	if r.span != nil {
		r.span.LogKV("value", rep)
	}
	return rep, nil
}

func (r conn) GetBytes(key string) ([]byte, error) {
	if r.span != nil {
		r.span.LogKV("command", "GetBytes", "key", key)
	}
	rep, err := redis.Bytes(r.conn.Do("Get", key))
	if err == redis.ErrNil {
		return nil, nil
	}
	if err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return nil, err
	}
	if r.span != nil {
		r.span.LogKV("value", string(rep))
	}
	return rep, nil
}

func (r conn) Del(key string) error {
	if r.span != nil {
		r.span.LogKV("command", "Del", "key", key)
	}
	if _, err := r.conn.Do("Del", key); err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return err
	}
	return nil
}

func (r conn) Lock(key string, timeout int) (bool, error) {
	if r.span != nil {
		r.span.LogKV("command", "Lock", "key", key, "timeout", timeout)
	}
	reply, err := r.conn.Do("Set", key, 1, "EX", timeout, "NX")
	if err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return false, err
	}
	if r.span != nil {
		r.span.LogKV("value", reply)
	}
	if reply != nil {
		return reply.(string) == "OK", nil
	}
	return false, nil
}

func (r conn) Unlock(key string) error {
	if r.span != nil {
		r.span.LogKV("command", "UnLock", "key", key)
	}
	return r.Del(key)
}

func (r conn) Incr(key string) (int64, error) {
	if r.span != nil {
		r.span.LogKV("command", "Incr", "key", key)
	}
	rep, err := r.conn.Do("Incr", key)
	if err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return 0, err
	}

	return rep.(int64), nil
}

func (r conn) Expire(key string, expire int) error {
	if r.span != nil {
		r.span.LogKV("command", "Expire", "key", key)
	}
	if _, err := r.conn.Do("Expire", key, expire); err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return err
	}

	return nil
}

func (r conn) LPush(key string, value interface{}) error {
	if r.span != nil {
		r.span.LogKV("command", "LPush", "key", key, "value", value)
	}
	if _, err := r.conn.Do("LPush", key, value); err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return err
	}

	return nil
}

func (r conn) LPop(key string) (interface{}, error) {
	if r.span != nil {
		r.span.LogKV("command", "LPop", "key", key)
	}
	rep, err := redis.String(r.conn.Do("LPop", key))
	if err == redis.ErrNil {
		return "", nil
	}
	if err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return "", err
	}
	if r.span != nil {
		r.span.LogKV("value", rep)
	}

	return rep, nil
}

func (r conn) RPush(key string, value interface{}) error {
	if r.span != nil {
		r.span.LogKV("command", "RPush", "key", key, "value", value)
	}
	if _, err := r.conn.Do("RPush", key, value); err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return err
	}

	return nil
}

func (r conn) RPop(key string) (string, error) {
	if r.span != nil {
		r.span.LogKV("command", "RPop", "key", key)
	}
	rep, err := redis.String(r.conn.Do("RPop", key))
	if err == redis.ErrNil {
		return "", nil
	}
	if err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return "", err
	}
	if r.span != nil {
		r.span.LogKV("value", rep)
	}

	return rep, nil
}

func (r conn) Do(commandName string, key string, args ...interface{}) (interface{}, error) {
	if r.span != nil {
		r.span.LogKV("cmd", commandName, "key", key, "args", args)
	}
	var doArgs []interface{}
	doArgs = append(doArgs, key, args)
	// 注意：若get找不到key，则 replay = nil；若要修改逻辑，需要同步修改mock逻辑
	replay, err := r.conn.Do(commandName, args...)
	if err != nil {
		if r.span != nil {
			r.span.SetTag("error", true)
		}
		return nil, err
	}
	if r.span != nil {
		r.span.LogKV("value", replay)
	}
	return replay, nil
}
