package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	for i := 1; i <= 3; i++ {
		id := i
		g.Go(func() error {
			if id == 2 {
				time.Sleep(1 * time.Second)
				fmt.Printf("Worker %d: I failed!\n", id)
				return fmt.Errorf("worker %d crashed", id)
			}
			fmt.Printf("Worker %d: Starting long task...\n", id)
			select {
			case <-time.After(5 * time.Second):
				fmt.Printf("Worker %d: Finished work\n", id)
				return nil
			case <-ctx.Done():
				fmt.Printf("Worker %d: Stopping (received cancel signal)\n", id)
				return ctx.Err()
			}
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("Main: Group failed with error: %v\n", err)
	}
}
