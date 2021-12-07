package main

import "fmt"

func main() {
	var i, j, k, counter int

	fmt.Println("Enter two integers")
	fmt.Scanf("%d %d", &i, &j)
	k = -1
	if i > j {
		k = i % j
	} else if j > i {
		k = j % i
	}
	for l := 1; l <= 10; l++ {
		if j%l == 0 && i%l == 0 {
			counter++
		}

	}
	if k == 0 {
		counter++
	}
	fmt.Printf("%d\n ", counter)
}
