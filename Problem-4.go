package main

import (
	"fmt"
)

func countMultiples(numbers []int) map[int]int {
	counts := make(map[int]int)

	for _, num := range numbers {
		for i := 1; i <= 9; i++ {
			if num%i == 0 {
				counts[i]++
			}
		}
	}

	return counts
}

func main() {
	numbers := []int{1, 2, 8, 9, 12, 46, 76, 82, 15, 20, 30}
	counts := countMultiples(numbers)

	fmt.Println(counts)
}
