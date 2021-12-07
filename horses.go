package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var testCases, length int
	var intLine string
	//var ar []int
	fmt.Scanf("%d", &testCases)
	fmt.Scanf("%d", &length)
	fmt.Scanf("%v", &intLine)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
	}
}
