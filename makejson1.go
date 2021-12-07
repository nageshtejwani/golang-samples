package main

import (
	"encoding/json"
	"fmt"
	"strings"
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

func main() {
	var personDetailsMap map[string]string
	var person Person
	var confirmation string = "Y"
	personDetailsMap = make(map[string]string)
	var personDetailsSlice = make([]map[string]string, 1)
	var index int = 0
	var jsonData []byte
	printProgramInformation()
	for confirmation != "X" {

		fmt.Println("Please enter your name::")
		_, err := fmt.Scan(&person.name)
		if err != nil {
			fmt.Println("Error occurred while reading name  ", err)
			return
		}
		fmt.Println("Please enter your address::")
		_, err = fmt.Scan(&person.address)
		if err != nil {
			fmt.Println("Error occurred  while reading address ", err)
			return
		}
		personDetailsMap["name"] = person.name
		personDetailsMap["address"] = person.address
		if index == 0 {
			personDetailsSlice[index] = personDetailsMap
		} else {
			personDetailsSlice = append(personDetailsSlice, personDetailsMap)
		}

		fmt.Println("Please enter X to discontinue,or any other key to continue")
		_, err = fmt.Scan(&confirmation)
		if err != nil {
			fmt.Println("Error occurred  while reading confirmation ", err)
			return
		}
		confirmation = strings.ToUpper(confirmation)
		index = index + 1
	}
	finalMap := make(map[string][]map[string]string)
	finalMap["finalList"] = personDetailsSlice
	if len(personDetailsSlice) > 1 {
		data, err := json.Marshal(finalMap)
		jsonData = data
		if err != nil {
			fmt.Println("Error occurred,failed to create Json ", err)
		}
	} else {
		data, err := json.Marshal(personDetailsMap)
		jsonData = data
		if err != nil {
			fmt.Println("Error occurred,failed to create Json ", err)
		}
	}

	fmt.Println(string(jsonData))
}
