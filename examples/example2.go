package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		time.Sleep(2 * time.Second)

		fmt.Println("Finished work 1")
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)

		fmt.Println("Finished work 2")
	})

	err := group.Wait(parallelizer.WithTimeout(time.Second))

	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
	fmt.Println()

	time.Sleep(2 * time.Second)
}
