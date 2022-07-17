package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() error {
		return errors.New("something went wrong")
	})

	group.Add(func() error {
		time.Sleep(10 * time.Second)
		return nil
	})

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
