package main

import (
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	files := []string{"main.go", "apple.js", "README.md"}
	for _, path := range files {
		p := path
		g.Go(func() error {
			fmt.Printf("Scanning for: %s...\n", p)
			_, err := os.Stat(p)
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf(" Stop! Error found: %v\n", err)
	} else {
		fmt.Println(" Success! All files are where they should be.")
	}
}
