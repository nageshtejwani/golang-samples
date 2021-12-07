package main

import (
	"fmt"
)

func printProgramInformation() {
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")
	fmt.Println("**********Program for Json Creation ***************************")
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")

}

type Person struct {
	name    string
	address string
}

// func main() {
// 	var personDetailsMap map[string]string = make(map[string]string)
// 	var person Person
// 	printProgramInformation()
// 	fmt.Println("Please enter your name::")
// 	_, err := fmt.Scan(&person.name)
// 	if err != nil {
// 		fmt.Println("Error occurred while reading name  ", err)
// 		return
// 	}
// 	fmt.Println("Please enter your address::")
// 	_, err = fmt.Scan(&person.address)
// 	if err != nil {
// 		fmt.Println("Error occurred  while reading address ", err)
// 		return
// 	}
// 	personDetailsMap["name"] = person.name
// 	personDetailsMap["address"] = person.address
// 	jsonData, err := json.Marshal(personDetailsMap)
// 	if err != nil {
// 		fmt.Println("Error occurred,failed to create Json ", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))
// }
