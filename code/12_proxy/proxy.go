package _12_proxy

import (
	"fmt"
	"sync"
	"time"
)

// UserRepository 定义真实对象与代理共同遵守的接口。
type UserRepository interface {
	FindByID(id int) (*User, error)
}

type User struct {
	ID   int
	Name string
}

// DBUserRepository 真实对象：模拟一次慢查询。
type DBUserRepository struct{}

func (d *DBUserRepository) FindByID(id int) (*User, error) {
	time.Sleep(100 * time.Millisecond)
	return &User{ID: id, Name: fmt.Sprintf("user-%d", id)}, nil
}

// CachedUserRepository 缓存代理：命中缓存直接返回，未命中才回源查询。
type CachedUserRepository struct {
	real  UserRepository
	mu    sync.RWMutex
	cache map[int]*User
}

func NewCachedUserRepository(real UserRepository) *CachedUserRepository {
	return &CachedUserRepository{real: real, cache: make(map[int]*User)}
}

func (c *CachedUserRepository) FindByID(id int) (*User, error) {
	c.mu.RLock()
	if u, ok := c.cache[id]; ok {
		c.mu.RUnlock()
		return u, nil
	}
	c.mu.RUnlock()

	u, err := c.real.FindByID(id)
	if err != nil {
		return nil, err
	}

	c.mu.Lock()
	c.cache[id] = u
	c.mu.Unlock()
	return u, nil
}
