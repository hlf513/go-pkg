package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	"github.com/redis/go-redis/v9"
	"sync"
)

type Casbiner interface {
	RedisWatcher(channel string) error
	CheckPermissionForUser(userId int64, permission ...any) (bool, error)

	AddPermissionForRole(roleId int64, permission ...string) error
	AddPermissionsForRole(roleId int64, permission ...[]string) error
	DeletePermissionForRole(roleId int64, permission ...string) error

	DeleteRole(roleIds []int64) error
	DeleteUser(userIds []int64) error

	AddRoleForUser(userId int64, roleIds []int64) error
	DeleteRoleForUser(userId int64, roleIds []int64) error
	DeleteRolesForUser(userId int64) error

	Enforcer() *casbin.Enforcer
}

var once sync.Once
var cas Casbiner

func Instance(opts ...Option) Casbiner {
	once.Do(func() {
		opt := newOptions(opts...)
		cas = &cb{opts: opt}
	})
	return cas
}

type cb struct {
	opts Options
	e    *casbin.Enforcer
}

func (c *cb) AddPermissionsForRole(roleId int64, permission ...[]string) error {
	if _, err := c.e.AddPermissionsForUser(
		fmt.Sprintf("%s%d", PrefixRoleID, roleId),
		permission...,
	); err != nil {
		return err
	}
	return nil
}

func (c *cb) DeleteRolesForUser(userId int64) error {
	if _, err := c.e.DeleteRolesForUser(
		fmt.Sprintf("%s%d", PrefixUserID, userId),
	); err != nil {
		return err
	}
	return nil
}

func (c *cb) DeleteRoleForUser(userId int64, roleIds []int64) error {
	for _, rid := range roleIds {
		if _, err := c.e.DeleteRoleForUser(
			fmt.Sprintf("%s%d", PrefixUserID, userId),
			fmt.Sprintf("%s%d", PrefixRoleID, rid),
		); err != nil {
			return err
		}
	}
	return nil
}

func (c *cb) DeletePermissionForRole(roleId int64, permission ...string) error {
	if _, err := c.e.DeletePermissionForUser(fmt.Sprintf("%s%d", PrefixRoleID, roleId), permission...); err != nil {
		return err
	}
	return nil
}

func (c *cb) Enforcer() *casbin.Enforcer {
	return c.e
}

func (c *cb) RedisWatcher(channel string) error {
	w, err := rediswatcher.NewWatcher(c.opts.RedisOptions.Addr, rediswatcher.WatcherOptions{
		Options: redis.Options{
			Network:  "tcp",
			Password: c.opts.RedisOptions.Password,
		},
		Channel: "/" + channel,
		// Only exists in test, generally be true
		IgnoreSelf: true,
	})
	if err != nil {
		return err
	}
	// Initialize the enforcer.
	m, err := model.NewModelFromString(defaultModel)
	if err != nil {
		return err
	}
	p, err := GormAdapter(c.opts.DBOptions)
	if err != nil {
		return err
	}
	if c.e, err = casbin.NewEnforcer(m, p); err != nil {
		return err
	}

	// Set the watcher for the enforcer.
	if err = c.e.SetWatcher(w); err != nil {
		return err
	}

	// use the default callback
	if err = w.SetUpdateCallback(rediswatcher.DefaultUpdateCallback(c.e)); err != nil {
		return err
	}

	// Update the policy to test the effect.
	// You should see "[casbin rules updated]" in the log.
	if err = c.e.SavePolicy(); err != nil {
		return err
	}

	return nil
}

func (c *cb) CheckPermissionForUser(userId int64, permission ...any) (bool, error) {
	var p []any
	p = append(p, fmt.Sprintf("%s%d", PrefixUserID, userId))
	p = append(p, permission...)
	return c.e.Enforce(p...)
}

func (c *cb) AddPermissionForRole(roleId int64, permission ...string) error {
	_, err := c.e.AddPermissionForUser(
		fmt.Sprintf("%s%d", PrefixRoleID, roleId),
		permission...)
	return err
}

func (c *cb) DeleteRole(roleIds []int64) error {
	for _, rid := range roleIds {
		if _, err := c.e.DeleteRole(fmt.Sprintf("%s%d", PrefixRoleID, rid)); err != nil {
			return err
		}
	}
	return nil
}

func (c *cb) DeleteUser(userIds []int64) error {
	for _, uid := range userIds {
		if _, err := c.e.DeleteUser(fmt.Sprintf("%s%d", PrefixUserID, uid)); err != nil {
			return err
		}
	}
	return nil
}

func (c *cb) AddRoleForUser(userId int64, roleIds []int64) error {
	for _, rid := range roleIds {
		if _, err := c.e.AddRoleForUser(
			fmt.Sprintf("%s%d", PrefixUserID, userId),
			fmt.Sprintf("%s%d", PrefixRoleID, rid),
		); err != nil {
			return err
		}
	}
	return nil
}
