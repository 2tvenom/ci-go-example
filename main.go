package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello! %d", sum(1, 2))
}

func sum(d1, d2 int) int {
	return d1 + d2
}
