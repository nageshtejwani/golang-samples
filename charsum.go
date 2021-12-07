package main

import (
	"fmt"
	"strings"
)

const (
	modifier = 96
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	var charsum int
	input = strings.ToLower(input)
	runes := []rune(input)
	for x := 0; x < len(runes); x++ {
		charsum = charsum + int(runes[x]) - modifier
	}
	fmt.Printf("%d", charsum)

}
