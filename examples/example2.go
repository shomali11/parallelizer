package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	func1 := func() {
		time.Sleep(time.Minute)

		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	}

	func2 := func() {
		time.Sleep(time.Minute)

		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	}

	runner := parallelizer.Runner{Timeout: time.Second}
	err := runner.Run(func1, func2)

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
