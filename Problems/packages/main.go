package main

import (
	"fmt"
	"github.com/walonCode/mathutils/mathutils"
)

func main(){
	xs:= []float64{1,2,3,4,5,7,10}
	max := mathutils.Max(xs)
	min := mathutils.Min(xs)
	fmt.Println("max is:",max, "min is: ",min)

	sum := mathutils.Sum(xs)
	fmt.Println("sum is:",sum)
}