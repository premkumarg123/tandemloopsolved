package main

import "fmt"

func generateSeries2(x int) {
	for i := 1; i <= x; i++ {
		fmt.Printf("%d ", 2*i-1)
	}
	fmt.Println()
}

func main() {
	input := 6
	if input%2 == 0 {
		generateSeries2(input - 1)
	} else {
		generateSeries2(input)
	}

}
