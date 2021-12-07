package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	runes := []rune(input)
	for x := 0; x < len(runes); x++ {
		if runes[x] >= 65 && runes[x] <= 90 {
			fmt.Printf("%s", strings.ToLower(string(runes[x])))
		} else if runes[x] >= 97 && runes[x] <= 122 {
			fmt.Printf("%s", strings.ToUpper(string(runes[x])))
		}
	}
}
