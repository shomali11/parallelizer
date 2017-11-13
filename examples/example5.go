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
		fmt.Print("Worker 1")
	})

	fmt.Println()
	fmt.Println("We did not wait!")

	time.Sleep(time.Second)
}
