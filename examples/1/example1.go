package main

import (
	"fmt"

	"github.com/shomali11/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() error {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
		return nil
	})

	group.Add(func() error {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
		return nil
	})

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
