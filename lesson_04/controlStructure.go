package main

//for loops
// if statement
//switch

import "fmt"

func main(){
	//for loop that checks for odd and even number
	for i:=1; i <= 10; i++ {
		if i % 2 == 1{
			fmt.Println(i, "odd" )
		}else{
			fmt.Println(i,"even")
		}	
	}
	// divisibleByThree()
	fizzBuzz()
	isOddOrEven(56)
}

func isOddOrEven(num float64){
	if int(num) % 2 == 0 {
		fmt.Println(num ,"Even");
	}else {
		fmt.Println(num, "Odd")
	}
}

func divisibleByThree(){
	for i:=1; i<=100; i++ {
		if i % 3 == 0{
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
			case i%3==0 && i%5==0: fmt.Println(i, "Fizzbuzz")
			case i%3 ==0:fmt.Println(i,"Fizz")
			case i%5 ==0:fmt.Println(i, "Buzz")
			default:fmt.Println(i)
		}
	}
}

