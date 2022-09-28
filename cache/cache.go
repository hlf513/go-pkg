package cache

import "time"

type Cache interface {
	Set(key, value string) error
	SetExpire(key, value string, timeout time.Duration) error
	Get(key string) error
}
