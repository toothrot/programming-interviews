package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"os"
	"math"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(b), "\n")
	line1 := []byte(lines[0])
	line2 := []byte(lines[1])

	line1map := make(map[byte]int)
	line2map := make(map[byte]int)
	changesmap := make(map[byte]float64)

	for _, b := range line1 {
		line1map[b] += 1
	}
	for _, b := range line2 {
		line2map[b] += 1
	}
	for k, v := range line1map {
		changesmap[k] = math.Abs(float64(v - line2map[k]))
	}
	for k, v := range line2map {
		_, ok := changesmap[k]
		if !ok {
			changesmap[k] = math.Abs(float64(v - line1map[k]))
		}
	}
	var sum int64
	for _, v := range changesmap {
		sum += int64(v)
	}
	fmt.Printf("%d", sum)
}
