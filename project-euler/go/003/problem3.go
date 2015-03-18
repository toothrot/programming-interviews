package main

import "fmt"
import "math"

func maxInt(numbers []int) int {
	max := 0
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func greatest_prime_factor(n int) int {
	factors := []int{1}
	floatn := float64(n)
	sqrt := math.Sqrt(floatn)

	for i := 2; float64(i) < sqrt; i++ {
		if n%i == 0 {
			if isPrime(i) {
				factors = append(factors, i)
			}
		}
	}

	return maxInt(factors)
}

func problem3(n int) int {
	return greatest_prime_factor(n)
}

func main() {
	result := problem3(600851475143)
	fmt.Println("Answer:", result)
}
