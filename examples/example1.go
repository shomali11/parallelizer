package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.DefaultGroup()

	group.Add(func() {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	})

	group.Add(func() {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	})

	err := group.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
