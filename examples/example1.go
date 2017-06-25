package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	runner := parallelizer.Runner{}

	runner.Add(func() {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	})

	runner.Add(func() {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	})

	err := runner.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
