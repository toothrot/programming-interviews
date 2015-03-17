package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return (fibonacci(n-1) + fibonacci(n-2))
	}
}

func problem2() int {
	sum := 0
	result := 0
	for i := 0; result <= 4000000; i++ {
		if result%2 == 0 {
			sum += result
		}
		result = fibonacci(i)
	}
	return sum
}

func main() {
	result := problem2()
	fmt.Println("Answer:", result)
}
