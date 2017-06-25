package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	func1 := func() {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	}

	func2 := func() {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	}

	runner := parallelizer.Runner{}
	hasFinished := runner.Run(func1, func2)

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Timed out? %t", !hasFinished)
}
