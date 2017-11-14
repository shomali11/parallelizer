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
		fmt.Println("Finished work")
	})

	fmt.Println("We did not wait!")

	time.Sleep(time.Second)
}
