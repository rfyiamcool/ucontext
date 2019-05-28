# ucontext

support multi value in context, and match golang std context

## Usage

```go
package main

import (
	"context"
	"time"

	"github.com/rfyiamcool/ucontext"
)

func main() {
	pctx := ucontext.New()
	pctx.Set("blog", "xiaorui.cc")
	pctx.Get("blog")

	ctx, cancel := context.WithCancel(pctx)
	cctx, _ := context.WithCancel(ctx)
	cancel()

	select {
	case <-cctx.Done():
	case <-time.After(1 * time.Second):
		panic("cancel failed")
	}
}
```