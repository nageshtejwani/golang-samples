package main

import "fmt"

func main() {
	var l, r, k int
	var counter int
	fmt.Scanf("%d %d %d", &l, &r, &k)
	for x := l; x <= r; x++ {
		if x%k == 0 {
			counter++
		}
	}
	fmt.Printf("%d \n", counter)
}
