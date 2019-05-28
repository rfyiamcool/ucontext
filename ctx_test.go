package ucontext

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	pctx, cancel := context.WithCancel(New())
	cancel()
	select {
	case <-pctx.Done():
	case <-time.After(1 * time.Second):
	}
}

func TestStore(t *testing.T) {
	pctx := New()
	pctx.Set("k", "v")
	value, ok := pctx.Get("k")
	if !ok {
		t.Error("not get")
	}
	assert.Equal(t, "v", value.(string))

	ctx, cancel := context.WithCancel(pctx.Context)
	cctx, _ := context.WithCancel(ctx)
	cancel()
	select {
	case <-cctx.Done():
	case <-time.After(1 * time.Second):
	}
}

func TestCancel(t *testing.T) {
	cctx, cancel := WithCancel(New())
	cctx.Set("key", "value")
	cancel()
	select {
	case <-cctx.Done():
	case <-time.After(1 * time.Second):
	}
}
