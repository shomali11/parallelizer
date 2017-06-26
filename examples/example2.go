package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
	"time"
)

func main() {
	runner := parallelizer.Runner{Timeout: time.Second}

	runner.Add(func() {
		time.Sleep(5 * time.Second)

		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	})

	runner.Add(func() {
		time.Sleep(5 * time.Second)

		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	})

	err := runner.Run()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
	fmt.Println()

	time.Sleep(10 * time.Second)
}
