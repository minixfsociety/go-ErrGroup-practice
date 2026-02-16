package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	fmt.Println("Manager: Starting tasks...")
	for i := 1; i <= 3; i++ {
		id := i
		g.Go(func() error {
			fmt.Printf(" Worker %d: Started working...\n", id)
			time.Sleep(1 * time.Second)
			if id == 2 {
				return fmt.Errorf("worker %d had a critical failure", id)
			}
			fmt.Printf("Worker %d: Finished successfully\n", id)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf(" Manager: Group failed with error: %v\n", err)
	} else {
		fmt.Println(" Manager: All tasks completed successfully!")
	}
}
