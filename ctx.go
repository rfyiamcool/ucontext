package ucontext

import (
	"context"
	"time"
)

type Context struct {
	context.Context
	Store map[string]interface{}
	// more...
}

// New a context
func New() *Context {
	return &Context{
		Context: context.Background(),
		Store:   make(map[string]interface{}, 10),
	}
}

func NewWithCtx(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
		Store:   make(map[string]interface{}, 10),
	}
}

func (c *Context) Set(k string, v interface{}) {
	c.Store[k] = v
}

func (c *Context) Get(k string) (interface{}, bool) {
	v, ok := c.Store[k]
	return v, ok
}

func (c *Context) Replace(m map[string]interface{}) {
	c.Store = m
}

// CancelFunc tells an operation to abandon its work
type CancelFunc context.CancelFunc

// WithCancel returns a copy of parent with a new Done channel
func WithCancel(parent *Context) (*Context, CancelFunc) {
	ctx := *parent
	child, cancel := context.WithCancel(parent.Context)
	ctx.Context = child
	return &ctx, CancelFunc(cancel)
}

// WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d
func WithDeadline(parent *Context, d time.Time) (*Context, CancelFunc) {
	ctx := *parent
	child, cancel := context.WithDeadline(parent.Context, d)
	ctx.Context = child
	return &ctx, CancelFunc(cancel)
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
func WithTimeout(parent *Context, timeout time.Duration) (*Context, CancelFunc) {
	ctx := *parent
	child, cancel := context.WithTimeout(parent.Context, timeout)
	ctx.Context = child
	return &ctx, CancelFunc(cancel)
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent *Context, key, val interface{}) *Context {
	ctx := *parent
	ctx.Context = context.WithValue(parent.Context, key, val)
	return &ctx
}
