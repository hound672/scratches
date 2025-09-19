package synctest

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

func TestWithDeadlineAfterDeadline(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		deadline := time.Now().Add(1 * time.Second)
		ctx, _ := context.WithDeadline(t.Context(), deadline)

		time.Sleep(time.Until(deadline))
		synctest.Wait()
		if err := ctx.Err(); err != context.DeadlineExceeded {
			t.Fatalf("context not canceled after deadline")
		}
	})
}
