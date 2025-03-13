package main

import "fmt";

func main(){
	var x[]float64
	fmt.Println(x)

	//using the make function to make a slice
	slice := make([]float64, 5)
	fmt.Println(slice)

	//make also allowed a third argument
	secondSlice:= make([]int, 5, 10)
	fmt.Println(secondSlice)

	//another way to make slice is to use the [low:high]
	arr := []float64{1,3,4,5,6}
	mySlice := arr[2:4]
	fmt.Println(mySlice)

	//using the append method
	slice1 := []int{1,2,3}
	slice2 := append(slice1,4,5)
	fmt.Println(slice1, slice2)

	//using the copy method copy(e1, e2)
	//e1 is the slice to be copied into and e2 is the slice or array copied from
	slice3 := make([]int, 2)
	copy(slice3,slice1)
	fmt.Println("slice 1",slice1, "slice 3: ",slice3)

	//excersie

	//learnt a new way of get the smallest number in a slice
	newSlice := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	//this methods work
	smallest := newSlice[0]

	for _, num := range newSlice{
		if num < smallest{
			smallest = num
		}
	}
	fmt.Println(smallest)


	//this doesn't
	smallNum := 0

	
	for _,num := range newSlice{
		if num <= smallNum{
			smallNum = num
		}
	}
	fmt.Println(smallNum)
}