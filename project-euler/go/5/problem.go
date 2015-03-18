package main

import "fmt"

func problem() int {
	found := 0

	for i := 0; ; i++ {
		for j := 1; j <= 20; j++ {
			found = i
			if i%j != 0 {
				found = 0
				break
			}
		}
		if found > 0 {
			break
		}
	}
	return found
}

func main() {
	result := problem()
	fmt.Println("Answer:", result)
}
