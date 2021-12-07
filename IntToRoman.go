package main

import "fmt"

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

func intToRoman(num int) string {
	//s := string(num)
	//runes := []rune(s)
	return "a"
}
func main() {
	//fmt.Printf("%s \n", intToRoman(10))
	var x int = 12
	var counter = 0
	for x > 10 {
		x = x % 10
		counter++
		fmt.Printf("%d \n", x)

	}
	switch counter {
	case 1 :
		fmt.Printf()

	}
}
