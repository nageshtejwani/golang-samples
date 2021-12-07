package main

import (
	"fmt"
	"os"
	"strings"
)

func printErrorMsg(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

func assembleLines(lines []string) []Person {
	var personDetailsSlice []Person
	for _, line := range lines {
		personDetails := strings.Fields(line)
		if len(personDetails) != 2 {
			fmt.Println("Incorrect file contents ,expecting only 2 words per line")
			os.Exit(1)
		}
		person := Person{fName: personDetails[0], lName: personDetails[1]}
		personDetailsSlice = append(personDetailsSlice, person)

	}
	return personDetailsSlice
}

func printResults(result []Person) {
	for _, person := range result {
		fmt.Println("***************************")
		fmt.Println("First Name  ::", person.fName)
		fmt.Println("Last  Name  ::", person.lName)
		fmt.Println("***************************")

	}
}

func readFileName() string {
	var fileName string
	fmt.Println("Enter a file name to be read with complete path :: ")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		printErrorMsg(err, "Error occurred while reading file name ")
		os.Exit(1)
	}
	return fileName
}

type Person struct {
	fName string
	lName string
}

// func main() {

	var text []string
	file, err := os.Open(readFileName())
	if err != nil {
		printErrorMsg(err, "Error occurred while reading file  ")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	result := assembleLines(text)
	fmt.Println("**** Printing the result ::")
	printResults(result)
}
