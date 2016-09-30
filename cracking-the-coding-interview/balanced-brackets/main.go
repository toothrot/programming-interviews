package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	closeMap = map[byte]byte{
		byte(')'): byte('('),
		byte('}'): byte('{'),
		byte(']'): byte('['),
	}
	OPENS  = []byte{'(', '{', '['}
	CLOSES = []byte{')', '}', ']'}
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(b), "\n")

	for i, line := range lines {
		if i == 0 {
			continue
		}
		if isBalanced([]byte(line)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func isBalanced(line []byte) bool {
	var lastOpen byte
	opens := []byte{}
	openCounts := map[byte]int{
		byte('('): 0,
		byte('{'): 0,
		byte('['): 0,
	}
	for _, b := range line {
		if isOpen(b) {
			opens = append(opens, b)
			openCounts[b] += 1
		} else if isClose(b) {
			if len(opens) == 0 {
				return false
			}
			openB, _ := closeMap[b]
			lastOpen, opens = opens[len(opens)-1], opens[:len(opens)-1]
			if lastOpen != openB {
				return false
			}
			if openCounts[openB] == 0 {
				return false
			}
			openCounts[openB] -= 1
		} else {
			return false
		}
	}
	for _, v := range openCounts {
		if v > 0 {
			return false
		}
	}
	return true
}

func isOpen(b byte) bool {
	for _, o := range OPENS {
		if b == o {
			return true
		}
	}
	return false
}

func isClose(b byte) bool {
	for _, o := range CLOSES {
		if b == o {
			return true
		}
	}
	return false
}
