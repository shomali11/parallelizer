package main

import (
	"fmt"
	"time"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup(parallelizer.WithPoolSize(1), parallelizer.WithJobQueueSize(1))
	defer group.Close()

	for i := 1; i <= 10; i++ {
		group.Add(func() {
			time.Sleep(time.Second)
		})

		fmt.Println("Job added at", time.Now().Format("04:05"))
	}

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
