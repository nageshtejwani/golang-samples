	package main

	import "fmt"

	const (
		motu     = "Motu"
		patlu    = "Patlu"
		modifier = 2
	)

	func main() {
		var x int
		var counter int = 1
		fmt.Scanf("%d", &x)
		for x >= 1 {
			x = x - counter
			if x < 0 {
				fmt.Printf("%s \n", patlu)
				break
			}
			x = x - 2*counter
			if x < 0 {
				fmt.Printf("%s \n", motu)
				break
			}
			counter++
		}

	}
