package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_POS = 2147483647
	MAX_NEG = -2147483648
)

func main() {
	var x int32 = 1534236469
	var resultString strings.Builder
	s := fmt.Sprint(x)
	var endCondn int = 0
	runes := []rune(s)
	k := len(runes)
	if x < 0 {
		resultString.WriteString("-")
		endCondn = 1
	}
	k = k - 1
	for m := k; m >= endCondn; m-- {
		resultString.WriteString(string(runes[m]))
	}
	r, err := strconv.Atoi(resultString.String())
	if err!= nil || r > MAX_POS || r < MAX_NEG {
		return 0
	}else{
		return r
	}
	// if err == nil {
	// 	return r
	// } else {
	// 	return -1
	// }
}
