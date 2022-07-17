package main

import (
	"fmt"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() error {
		fmt.Println("Worker 1")
		return nil
	})

	group.Add(func() error {
		fmt.Println("Worker 2")
		return nil
	})

	fmt.Println("Waiting for workers 1 and 2 to finish")

	group.Wait()

	fmt.Println("Workers 1 and 2 have finished")

	group.Add(func() error {
		fmt.Println("Worker 3")
		return nil
	})

	fmt.Println("Waiting for worker 3 to finish")

	group.Wait()

	fmt.Println("Worker 3 has finished")
}
