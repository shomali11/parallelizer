package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	options := &parallelizer.Options{Timeout: time.Second}
	group := parallelizer.NewGroup(options)

	group.Add(func() {
		time.Sleep(2 * time.Second)

		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	})

	group.Add(func() {
		time.Sleep(2 * time.Second)

		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	})

	err := group.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
	fmt.Println()

	time.Sleep(3 * time.Second)
}
