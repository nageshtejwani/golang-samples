package main

import (
	"fmt"
	"strings"
)

const (
	cow   = "cow"
	bird  = "bird"
	snake = "snake"
	EAT   = "EAT"
	MOVE  = "MOVE"
	SPEAK = "SPEAK"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func assembleAnimalData() map[string]Animal {
	animalDetails := make(map[string]Animal, 3)
	animalDetails[cow] = Animal{

		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	animalDetails[bird] = Animal{

		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	animalDetails[snake] = Animal{

		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	return animalDetails
}

func programInformation() {
	fmt.Println("*************************************")
	fmt.Println("Program prints animals information ::")
	fmt.Println("*************************************")
	fmt.Println("*************************************")
	fmt.Println("******sample input >cow eat *********")
	fmt.Println("******sample input  >cow move ********")
	fmt.Println("******sample input  >cow speak *******")
}
func (animal Animal) Eat() {
	fmt.Printf("This animal eats %s", animal.food)
}
func (animal Animal) Move() {
	fmt.Printf("This animal movement is %s", animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Printf("This animal sound is  %s", animal.noise)
}

// func PrintErrorMsg(err error, msg string) {
// 	if err != nil {
// 		fmt.Println(msg, err)
// 	}
// }

// func TakeInput(msg string, noOfParams int) []string {
// 	result := make([]string, 2)
// 	fmt.Print(msg)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Split(bufio.ScanWords)
// 	scanner.Scan()
// 	switch noOfParams {
// 	case 1:
// 		result[0] = scanner.Text()
// 	case 2:
// 		result[0] = scanner.Text()
// 		scanner.Scan()
// 		result[1] = scanner.Text()
// 	default:
// 		fmt.Println("Incorrect number of parameters")
// 		os.Exit(1)
// 	}
// 	return result
// }
// func PrintAnimalAttribute(animal Animal, attribute string) {

// 	var info = strings.ToUpper(attribute)
// 	switch info {
// 	case EAT:
// 		animal.Eat()
// 	case MOVE:
// 		animal.Move()
// 	case SPEAK:
// 		animal.Speak()
// 	default:
// 		fmt.Println("Incorrect parameters have been passed ")
// 	}

// }

// func getAnimalInformation(animal string, attribute string, animalInformation map[string]Animal) {
// 	animal = strings.ToLower(animal)
// 	attribute = strings.ToLower(attribute)
// 	switch animal {
// 	case cow:
// 		cowVar, _ := animalInformation[cow]
// 		PrintAnimalAttribute(cowVar, attribute)
// 	case bird:
// 		birdVar, _ := animalInformation[bird]
// 		PrintAnimalAttribute(birdVar, attribute)
// 	case snake:
// 		snakeVar, _ := animalInformation[snake]
// 		PrintAnimalAttribute(snakeVar, attribute)
// 	default:
// 		fmt.Println("Incorrect parameters have been passed ")
// 	}
// }

//func main() {
// 	var confirmation string = "Y"
// 	programInformation()
// 	animalInfo := assembleAnimalData()
// 	for confirmation == "Y" {
// 		fmt.Println()
// 		inputLine := TakeInput("Enter text  >", 2)
// 		getAnimalInformation(inputLine[0], inputLine[1], animalInfo)
// 		fmt.Println()
// 		inputLine = TakeInput("Enter Y or y to continue or any other key to discontinue  ::>", 1)
// 		confirmation = strings.ToUpper(inputLine[0])
// 	}
//}
