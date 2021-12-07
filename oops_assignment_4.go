package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	cow       = "cow"
	bird      = "bird"
	snake     = "snake"
	newanimal = "newanimal"
	query     = "query"
	EAT       = "EAT"
	MOVE      = "MOVE"
	SPEAK     = "SPEAK"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (cow Cow) Eat() {
	fmt.Printf("This animal eats %s", cow.food)
}
func (cow Cow) Move() {
	fmt.Printf("This animal movement is %s", cow.locomotion)

}
func (cow Cow) Speak() {
	fmt.Printf("This animal sound is  %s", cow.noise)
}

func (bird Bird) Eat() {
	fmt.Printf("This animal eats %s", bird.food)
}
func (bird Bird) Move() {
	fmt.Printf("This animal movement is %s", bird.locomotion)

}
func (bird Bird) Speak() {
	fmt.Printf("This animal sound is  %s", bird.noise)
}

func (snake Snake) Eat() {
	fmt.Printf("This animal eats %s", snake.food)
}
func (snake Snake) Move() {
	fmt.Printf("This animal movement is %s", snake.locomotion)

}
func (snake Snake) Speak() {
	fmt.Printf("This animal sound is  %s", snake.noise)
}

func programInformation() {
	fmt.Println("*************************************")
	fmt.Println("Program prints/create animals information ::")
	fmt.Println("*************************************")
	fmt.Println("*************************************")
	fmt.Println("******sample input >newanimal '<str>' cow *******")
	fmt.Println("******sample input  >query '<str>'  ********")
}

func assembleAnimalData() map[string]Animal {
	animalDetails := make(map[string]Animal, 3)
	cowAnimal := Cow{

		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	birdAnimal := Bird{

		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snakeAnimal := Snake{

		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
	animalDetails[cow] = cowAnimal
	animalDetails[bird] = birdAnimal
	animalDetails[snake] = snakeAnimal

	return animalDetails
}

func TakeInput(msg string, noOfParams int) []string {
	fmt.Printf(msg)
	result := make([]string, 3)
	var counter int = 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for counter < noOfParams {
		scanner.Scan()
		result[counter] = scanner.Text()
		counter++
	}
	return result
}

func getAnimalInformation(animal string, attribute string, animalInformation map[string]Animal) {
	_, result := animalInformation[animal]
	if !result {
		println("Animal not found in the list")
		return
	}
	attribute = strings.ToUpper(attribute)
	switch attribute {
	case EAT:
		animalInformation[animal].Eat()
	case MOVE:
		animalInformation[animal].Move()
	case SPEAK:
		animalInformation[animal].Speak()
	default:
		fmt.Println("Incorrect parameters have been passed ")
	}
}

func addAnimal(name string, animalType string, animalMap map[string]Animal) {
	animalType = strings.ToLower(animalType)
	switch animalType {
	case cow:
		animalMap[name] = animalMap[cow]
	case bird:
		animalMap[name] = animalMap[bird]
	case snake:
		animalMap[name] = animalMap[snake]
	default:
		fmt.Println("Animal type not found")
	}

}

// func main() {
// 	programInformation()
// 	animalDetails := assembleAnimalData()
// 	var confirmation string = "Y"
// 	for confirmation == "Y" {
// 		fmt.Println()
// 		inputLine := TakeInput("Enter text  >", 3)
// 		queryType := strings.ToLower(inputLine[0])
// 		switch queryType {
// 		case newanimal:
// 			addAnimal(inputLine[1], inputLine[2], animalDetails)
// 		case query:
// 			getAnimalInformation(inputLine[1], inputLine[2], animalDetails)
// 		}
// 		fmt.Println()
// 		inputLine = TakeInput("Enter Y or y to continue or any other key to discontinue  ::>", 1)
// 		confirmation = strings.ToUpper(inputLine[0])
// 	}
// }
