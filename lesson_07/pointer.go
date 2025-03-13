package main

import "fmt"

func main(){
	x := 1.5
	square(&x)
	fmt.Println(x)
	
	y:=5;
	zero(&y)
	fmt.Println(y)

	xPtr := new(int)
	five(xPtr)
	fmt.Println(*xPtr)

}

func zero(xPtr *int){
	*xPtr = 0
}

func square(x *float64){
	*x = *x * *x
}

func five(xPtr *int){
	*xPtr = 5
}

// * points to the data in the memory of a given location
// & returns the address of a variable
// new function returns a pointer