package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	var zCounter int
	var oCounter int
	input = strings.ToUpper(input)
	runes := []rune(input)
	if strings.HasPrefix(input, "Z") {
		z := strings.LastIndex(input, "Z")
		o := strings.Index(input, "O")
		for x := 0; x <= z; x++ {
			if runes[x] == rune('Z') {
				zCounter++
			} else {
				fmt.Println("No")
				return
			}
		}
		for o < len(input) {
			if runes[o] == rune('O') {
				oCounter++
			} else {
				fmt.Println("No")
				return
			}
			o++
		}

	} else {
		fmt.Println("No")
	}

	if 2*zCounter == oCounter {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
