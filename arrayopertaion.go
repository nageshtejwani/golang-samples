package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	ONE   = "ONE"
	TWO   = "TWO"
	THREE = "THREE"
	FOUR  = "FOUR"
)

func PrintErrorMsg(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
func TakeInputSize() int {
	var input string
	fmt.Println("Enter size of a slice as an integer")
	_, err := fmt.Scan(&input)
	if err != nil {
		PrintErrorMsg(err, "Error occurred while reading integer ")
		os.Exit(1)
	}
	k, err := strconv.Atoi(input)
	if err != nil {
		PrintErrorMsg(err, "Error occurred while input conversion ")
		os.Exit(1)
	}
	return k
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

func split(intList []int) map[string][]int {
	var sliceLen int = len(intList)
	var sliceMap map[string][]int
	sliceMap = make(map[string][]int, 4)
	if sliceLen == 0 {
		fmt.Println("Input list is empty ")
		os.Exit(1)
	} else {
		k := sliceLen / 4
		slice1 := intList[0:k]
		slice2 := intList[k : 2*k]
		slice3 := intList[2*k : 3*k]
		slice4 := intList[3*k : sliceLen]
		sliceMap[ONE] = slice1
		sliceMap[TWO] = slice2
		sliceMap[THREE] = slice3
		sliceMap[FOUR] = slice4

	}
	return sliceMap
}

func sortIntSlice(c chan map[string][]int, keyElement string) {
	fmt.Printf("Thread Started with Key %s\n", keyElement)
	var sliceMap map[string][]int
	sliceMap = <-c
	tempSlice := sliceMap[keyElement]
	sort.Ints(tempSlice)
	fmt.Printf("Sorted Slice %v\n", tempSlice)
	sliceMap[keyElement] = tempSlice
}

// func main() {
// 	var intList []int
// 	var length int = TakeInputSize()
// 	intList = make([]int, length)
// 	var sliceMap map[string][]int
// 	var mergedList []int = make([]int, 0)
// 	TakeInput(intList)
// 	sliceMap = split(intList)
// 	fmt.Println(intList)
// 	c := make(chan map[string][]int)
// 	for key, _ := range sliceMap {
// 		go sortIntSlice(c, key)
// 		c <- sliceMap
// 	}
// 	for key, _ := range sliceMap {
// 		mergedList = append(mergedList, sliceMap[key]...)
// 	}
// 	sort.Ints(mergedList)
// 	fmt.Printf("Sorted List  %v\n", mergedList)

// }
