package main 

import "fmt";

func main(){

	// array in go is a collection of a single element type
	var x [5]int
	x[4] = 100;
	// fmt.Println(x[4])

	var Y [5]int;
	Y[0] = 78;
	Y[1] = 50;
	Y[2] = 80;
	Y[3] = 10;
	Y[4] = 10;

	var total int = 0
	//better way and concise way of using an for loop
	for i:= range len(Y) {
		total += Y[i]
	}
	// fmt.Println(float64(total / 5))

	//this is a slice instead of an array in go, arrays have fixed length
	var array []int;
	array = append(array, 24,50,70,50)
	// fmt.Println(array)

	var getTotal int
	// another clean way of write a good for loop
	for _, value := range array{
		getTotal += value;
	}
	// fmt.Println(getTotal)


	//shorter way of creating an array in go
	myArr := [10]float64{
		98,
		70,
		89,
		100,
		// 50.67,
	}

	//Note: The length of an array if defined will not change as go doesn't allowed that

	fmt.Println(len(myArr))
}



