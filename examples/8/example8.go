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

	fmt.Println("Waiting for workers 1 and 2 to finish")

	group.Wait()

	fmt.Println("Workers 1 and 2 have finished")

	group.Add(func() {
		fmt.Println("Worker 3")
	})

	fmt.Println("Waiting for worker 3 to finish")

	group.Wait()

	fmt.Println("Worker 3 has finished")
}
