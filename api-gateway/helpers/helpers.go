package helpers

import (
	"context"
	"time"
)

func NewCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 2*time.Second)
}
