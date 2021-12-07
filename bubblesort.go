package main

import (
	"fmt"
	"os"
	"strconv"
)

func PrintProgramInformation() {
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")
	fmt.Println("**********Program for Bubble Sort with 10 integers************")
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")

}
func PrintErrorMsg(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

func Swap(intList []int, k int) {
	var temp int
	temp = intList[k]
	intList[k] = intList[k+1]
	intList[k+1] = temp
}

func BubbleSort(intList []int) {
	var length int = len(intList)
	for x := 0; x < length-1; x++ {
		var isSwapped bool = false
		for y := 0; y < length-1; y++ {
			if intList[y] > intList[y+1] {
				Swap(intList, y)
				isSwapped = true
			}
		}
		if isSwapped == false {
			return
		}
	}
}

func TakeInput(intList []int) {
	var input string
	var counter int = 0
	for range intList {
		fmt.Println("Enter an Integer ")
		_, err := fmt.Scan(&input)
		if err != nil {
			PrintErrorMsg(err, "Error occurred while reading integer ")
			os.Exit(1)
		}
		intList[counter], err = strconv.Atoi(input)
		if err != nil {
			PrintErrorMsg(err, "Error occurred while input conversion ")
			os.Exit(1)
		}
		counter = counter + 1
	}
}

// func main() {
// 	var intList []int
// 	var length int = 10
// 	intList = make([]int, length)
// 	PrintProgramInformation()
// 	TakeInput(intList)
// 	fmt.Println("Input List ", intList)
// 	BubbleSort(intList)
// 	fmt.Println("Sorted List ", intList)
// }
