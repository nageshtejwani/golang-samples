package main

import (
	"fmt"
	"strings"
)

func printErrorMsg(k int, errMsg string) {
	fmt.Println("Number of items received ::  ", k)
	fmt.Println("Error occurred ", errMsg)
}

func checkString(str string) {
	strlength := len(str)
	cmpLength := strlength - 1
	if strings.IndexRune(str, 'I') == 0 && strings.LastIndex(str, "N") == cmpLength &&
		strings.IndexRune(str, 'A') != -1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

// func main() {
// 	var input string
// 	fmt.Println("Enter a string to find 'i'(start index) ,'a'(anywhere in the string) and 'n(end index)' letters in a string")
// 	fmt.Println("program is case-insensitive")
// 	i, err := fmt.Scan(&input)
// 	if err != nil {
// 		printErrorMsg(i, fmt.Sprint(err))
// 		return
// 	}
// 	modifiedstring := strings.ToUpper(input)
// 	checkString(modifiedstring)
// }
