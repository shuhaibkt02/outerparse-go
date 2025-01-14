package main

import (
	"flag"
	"fmt"
)

func main() {
	count := flag.Int("c", 1, "Number of times")
	name := flag.String("n", "Guest", "Name")
	message := flag.String("m", "Hello", "Message to display")

	flag.Parse()

	for i := 0; i < *count; i++ {
		fmt.Printf("%s: %s\n", *name, *message)
	}
}
