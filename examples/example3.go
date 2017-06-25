package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	func1 := func() {
		fmt.Println("func1")
	}

	func2 := func() {
		fmt.Println("func2")
	}

	runner := parallelizer.Runner{}
	err := runner.Run(func1, func2)

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
