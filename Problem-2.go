package main

import (
	"fmt"
)

func generateSeries(x int) {
	for i := 1; i <= x; i++ {
		fmt.Printf("%d ", 2*i-1)
	}
	fmt.Println()
}

func main() {
	input := 3
	generateSeries(input)
}
