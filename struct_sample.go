package main

import "fmt"

type Student struct {
	id   int
	name string
}

func main() {
	s1 := Student{id: 1, name: "Elon Musk"}
	s2 := Student{id: 2, name: "Nikola Tesla"}
	fmt.Println(s1, s2)
}
