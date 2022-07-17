package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() error {
		time.Sleep(2 * time.Second)

		fmt.Println("Finished work 1")

		return nil
	})

	group.Add(func() error {
		time.Sleep(2 * time.Second)

		fmt.Println("Finished work 2")

		return nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := group.Wait(parallelizer.WithContext(ctx))

	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
	fmt.Println()

	time.Sleep(2 * time.Second)
}
