package main

import (
	"fmt"
)

func printProgramInformation() {
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")
	fmt.Println("**********Program for Float to Integer Conversion************")
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")

}

func printErrorMsg(k int, errMsg string) {
	fmt.Println("Number of items received ::  ", k)
	fmt.Println("Error occurred ", errMsg)
}

// func main() {
// 	var x float64
// 	var y int64
// 	var conf string = "Y"
// 	var input string
// 	printProgramInformation()
// 	for conf == "Y" {
// 		fmt.Printf("Type a float number ::")
// 		i, err := fmt.Scan(&input)
// 		x, err = strconv.ParseFloat(input, 64)
// 		if err != nil {
// 			printErrorMsg(i, fmt.Sprint(err))
// 			return
// 		} else if math.MaxInt64 < x {
// 			y = int64(x)
// 			fmt.Println("The float number is higher than the max int64 value",
// 				" , truncated integer Value is : ", y)
// 		} else {
// 			y = int64(x)
// 			fmt.Printf(" Truncated integer Value is %d : ", y)
// 		}
// 		fmt.Println("Press Y to continue,else type any other key to exit .. ")
// 		j, err := fmt.Scan(&conf)
// 		if err != nil {
// 			printErrorMsg(j, fmt.Sprint(err))
// 		}
// 	}
// }
