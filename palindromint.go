package main

import "fmt"

const (
	MAX_POS = 2147483647
	MAX_NEG = -2147483648
)

func isPalindrome(x int) bool {
	var reverse int = 0
	var y int = x
	if x < 0 || x > MAX_POS || x < MAX_NEG || x == 10 {
		return false
	}
	if x > 0 && x <= 9 {
		return true
	}
	for x != 0 {
		reverse = reverse * 10
		reverse = reverse + x%10
		x = (x / 10)
	}
	if reverse == y {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(10))
}
