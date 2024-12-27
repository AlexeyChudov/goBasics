package context_pckg

import (
	"context"
	"fmt"
	"time"
)

func MainFunc() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		fmt.Scan()
		cancel()

	}()

	ProcessLongTask(ctx)

}

func ProcessLongTask(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Task Done")
	case <-ctx.Done():
		fmt.Println("Task canceled")
	}

}
