package examples

import (
	"context"
	"fmt"
	"time"
)

func RunContextExample() {
	ctx := context.Background()
	withValue(ctx)
}

func timeout(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("API called")
	case <-ctxWithTimeout.Done():
		fmt.Println("Timeout: ", ctxWithTimeout.Err())
	}

	fmt.Println("Done")
}

func withValue(ctx context.Context) {
	ctxValue := context.WithValue(ctx, "key", 123)
	fmt.Println(ctxValue.Value("key"))
}
