package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 5, 0)
	fmt.Println(total)

	numbers := []int{10, 10, 5, 5}
	fmt.Println(sumAll(numbers...))
}
