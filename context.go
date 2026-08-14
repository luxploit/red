package red

import "sync"

type Context struct {
	data map[any]any
	mu   sync.RWMutex
}

type ContextKey string

func contextProvider() (*Context, error) {
	return &Context{
		data: make(map[any]any),
	}, nil
}

func (c *Context) Set(key, val any) *Context {
	if key == nil {
		panic("red: context key cannot be nil")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = val
	return c
}

func (c *Context) Get(key any) any {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data[key]
}
