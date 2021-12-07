package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printProgramInformation() {
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")
	fmt.Println("**********Program for Slice Usage****************************")
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")

}

func printErrorMsg(k int, errMsg string) {
	fmt.Println("Number of items received ::  ", k)
	fmt.Println("Error occurred ", errMsg)
}

func validateInput(err error, input string, i int) int {
	switch {
	case err != nil:
		printErrorMsg(i, fmt.Sprint(err))
		return -1
	case input == "X" || input == "x":
		fmt.Println("Exiting the program ,bye !")
		return -1
	}
	return 0
}

func parseInput(input string) (int, int) {
	x, err := strconv.Atoi(input)
	if err != nil {
		printErrorMsg(x, fmt.Sprint(err))
		return -1, -1
	}
	return x, 0
}

func manageIntegerList(counter int, integersList []int, x int) (int, []int) {
	if counter <= 2 {
		index := sort.SearchInts(integersList, 0)
		integersList[index] = x
	} else {
		integersList = append(integersList, x)
	}
	counter = counter + 1
	fmt.Println()
	sort.Ints(integersList)
	fmt.Printf("Sorted Slice  :: %d", integersList)
	fmt.Println()
	return counter, integersList
}

// func main() {
// 	var input string
// 	var integersList = make([]int, 3)
// 	var counter int = 0
// 	var parsedValue int
// 	var statusCode int
// 	printProgramInformation()
// 	for {
// 		fmt.Println("Type an Integer or type X or x to discontinue ::")
// 		i, err := fmt.Scan(&input)
// 		if validateInput(err, input, i) == -1 {
// 			return
// 		}
// 		if parsedValue, statusCode = parseInput(input); statusCode == -1 {
// 			return
// 		}
// 		fmt.Println()
// 		fmt.Printf("Length and Capacity %d ,%d of a slice ", len(integersList), cap(integersList))
// 		counter, integersList = manageIntegerList(counter, integersList, parsedValue)
// 	}
}
