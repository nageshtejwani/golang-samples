package main

import (
	"fmt"
)

func increment(a *int) {
	fmt.Println()
	fmt.Printf("Before increment %d", *a)
	fmt.Println()
	*a++
	fmt.Printf("After increment %d", *a)
	fmt.Println()
}

func decrement(a *int) {
	fmt.Println()
	fmt.Printf("Before decrement %d", *a)
	fmt.Println()
	*a--
	fmt.Printf("After decrement %d", *a)
	fmt.Println()
}

// func main() {
// 	var i int = 1
// 	for i <= 1 {
// 		go increment(&i)
// 		go decrement(&i)
// 	}
// 	time.Sleep(5 * time.Second)
// }
