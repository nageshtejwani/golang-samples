package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	acceleration     = "Acceleration"
	velocity         = "Velocity"
	initDisplacement = "Initial Displacement"
	time             = "Time"
)

type Displacement struct {
	time     float64
	acce     float64
	velocity float64
	initDisp float64
}

func PrintProgramInformation() {
	fmt.Println("*************************************************************")
	fmt.Println("*************************************************************")
	fmt.Println("********** Dispacement calculation **************************")
	fmt.Println("******** Formula :::result=½ a *t**2 + v*t + s*****************")
	fmt.Println("*************************************************************")

}
func PrintErrorMsg(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

func PrintAllValues(disp Displacement) {
	fmt.Printf("Acceleration value  :: %f", disp.acce)
	fmt.Println()
	fmt.Printf("Velocity value ::  %f", disp.velocity)
	fmt.Println()
	fmt.Printf("Initial Displacement value  :: %f", disp.initDisp)
	fmt.Println()
}
func GenDisplaceFn(acce, velocity, disp float64) func(t float64) float64 {
	return func(time float64) float64 {
		return 0.5*acce*math.Pow(time, 2) + velocity*time + disp
	}
}

func ConvertInputToFloat(input string, inputType string) float64 {
	convertedValue, err := strconv.ParseFloat(input, 64)
	if err != nil {
		PrintErrorMsg(err, "Error occurred while converting  "+inputType)
		os.Exit(1)
	}
	return convertedValue
}
func ReadFloatValue(inputType string, disp *Displacement) {
	var input string
	fmt.Println("Enter " + inputType + " value as float  ")
	_, err := fmt.Scan(&input)
	if err != nil {
		PrintErrorMsg(err, "Error occurred while reading "+inputType+" value ")
		os.Exit(1)
	}
	switch inputType {
	case acceleration:
		disp.acce = ConvertInputToFloat(input, inputType)
	case velocity:
		disp.velocity = ConvertInputToFloat(input, inputType)
	case initDisplacement:
		disp.initDisp = ConvertInputToFloat(input, inputType)
	case time:
		disp.time = ConvertInputToFloat(input, inputType)
	}

}

func TakeInput(disp *Displacement) {
	ReadFloatValue(acceleration, disp)
	ReadFloatValue(velocity, disp)
	ReadFloatValue(initDisplacement, disp)
}

// func main() {
// 	PrintProgramInformation()
// 	var disp Displacement
// 	TakeInput(&disp)
// 	PrintAllValues(disp)
// 	fn := GenDisplaceFn(disp.acce, disp.velocity, disp.initDisp)
// 	ReadFloatValue(time, &disp)
// 	fmt.Println(fn(disp.time))

// }
