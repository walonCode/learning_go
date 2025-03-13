package main

//working with data types in Go
//Note: go only add values of the save type
//working with int and float for number
//strings for text
//boolean
//This are the basic types in go

import "fmt"

func main(){
	//explict types
	const intValue int32 = 50;
	const floatValue float32 = 76.00
	fmt.Print(int(intValue) + int(floatValue) ,"\n")

	//inferred types
	newIntValue := 78;
	newFloatValue := 80.9
	fmt.Println(newIntValue + int(newFloatValue))

	//working with string
	const name string = "WALON";
	newName := "Nama"
	fmt.Println("Walon"[1])
	fmt.Println(len(name))
	fmt.Println(newName)

	//boolean
	const isOpened bool = true;
	fmt.Println(!isOpened)
	fmt.Println((true && false) || (false && true) || !(false && true))

	//math operation in go
	fmt.Println(200 * 50)
	
}
