package main

import (
	"fmt"
)

func main(){
	arr := []float64{
		24.5,
		34.5,
		55.6,
	}
	var arr2 []float64
	answer,error := average(arr2)
	fmt.Println(answer,error)

	xs := []int{1,2,3}
	
	fmt.Println(average(arr))

	fmt.Println(add(xs...))

	//closure creating a function inside a function
	X := 0
	increment := func()int {
		X++
		return X
	}
	fmt.Println(increment())

	// nextEven := makeEvenGenerator()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())

	// fmt.Println(factorial(3))
	defer second()
	first()

}

func average(xs []float64)(float64,string) {
	total := 0.0
	error:=""
	if len(xs)==0{
		error = "Error occured"
		total = 0.0
	}else{
		error = "false"
	}
	for _, num:=range xs {
		total += num
	}
	average := total/float64(len(xs))
	return float64(average),error
}

//variadic function
func add(args ...int) int {
	total := 0;
	for _, v:= range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint{
	i:=uint(0)
	return func() (ret uint) {
		ret = i
		i+=2
		return
	}
}

func factorial(x uint)uint{
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first(){
	fmt.Println("1st")
}

func second(){
	fmt.Println("2nd")
}



