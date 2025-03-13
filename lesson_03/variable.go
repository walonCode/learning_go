package main

//working with variable
//also mada use of functions

import (
	"fmt"
)

func main(){
	//declaring and assigning the variable
	var text string = "Walon";

	//declaring and assigning later
	var name string
	name = "jalloh"
	fmt.Println(text + name)

	//excersie concatenation of string literals
	var y string
	y = "first"
	fmt.Println(y)
	y += "second"
	fmt.Println(y)
	y += "third"
	fmt.Println(y)

	//boolean check of string
	var middleName = "Walon"
	var lastName = "Jalloh"
	fmt.Println(middleName != lastName)

	//infered types
	w := 90
	fmt.Println(w == 1200)
	// multipleVariableDefinition()
	// doubler()
	// fahrenhietToCelsius()
	feetToMeter()
}

func multipleVariableDefinition(){
	const (
		a=5
		b=20
		c=50
	)
	fmt.Print(a,b,c ,"\n")
}

func doubler(){
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f",&input)

	output := input * 2;
	fmt.Println(output)
}

func fahrenhietToCelsius(){
	var value float64
	fmt.Println("Enter temperature in Fahrenhiet: ")
	fmt.Scanf("%f",&value)
	output := (value - 32) * 5/9
	fmt.Println(output)
}

func feetToMeter(){
	var value float64
	fmt.Println("Welcome to feet to meter converter")
	fmt.Println("Enter distance in feet: ")
	fmt.Scanf("%f",&value)
	output := value * 0.3048
	fmt.Println(output)
}