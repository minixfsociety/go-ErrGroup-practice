package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	var totalSize int64
	list := []string{"google.com", "youtube.com", "yandex.ru", "temu.com"}
	for _, url := range list {
		u := url
		g.Go(func() error {
			fmt.Println("Download:", u)
			time.Sleep(2 * time.Second)
			atomic.AddInt64(&totalSize, 100)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Good")
}
