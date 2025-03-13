package main

import "fmt"

func main(){
	value1 := 5
	value2 := 10
	swap(&value1, &value2)
	fmt.Println("memory swapping",value1, value2)

	normalswap(value1, value2)
	fmt.Println("normal swapping", value1, value2)

}

func swap(value1 *int, value2 *int){
	*value1, *value2 = *value2, *value1
}

func normalswap(value1 int, value2 int){
	temp:= value1;
	value1 = value2
	value2 = temp
}

// go has the same swapping technics as python