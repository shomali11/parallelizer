package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	options := &parallelizer.Options{WorkerPoolSize: 10}
	group := parallelizer.NewGroup(options)

	for i := 1; i <= 10; i++ {
		i := i
		group.Add(func() {
			fmt.Print(i, " ")
		})
	}

	err := group.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
