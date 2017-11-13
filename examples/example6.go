package main

import (
	"fmt"
	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		fmt.Println("Worker 1")
	})

	group.Add(func() {
		fmt.Println("Worker 2")
	})

	group.Wait()

	fmt.Println("Workers 1 and 2 have finished")

	group.Add(func() {
		fmt.Println("Worker 3")
	})

	group.Wait()

	fmt.Println("Worker 3 has finished")
}
