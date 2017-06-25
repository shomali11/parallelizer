package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	func1 := func() {
		fmt.Println("func1")
	}

	func2 := func() {
		fmt.Println("func2")
	}

	functions := []func(){func1, func2}

	runner := parallelizer.Runner{Timeout: time.Second}
	err := runner.Run(functions...)

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
