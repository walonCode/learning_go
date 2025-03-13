package main

import "fmt"

func main(){
	// arr:= []float64{
	// 	10,20,30,
	// }
	// // var myarr [5]float64
	// fmt.Println(add(arr))
	// fmt.Println(half(2))
	// arr := []float64{
	// 	1,2,3,4,5,
	// }
	// fmt.Println(highestNumber(arr...))
	// oddValue := makeOddGenerator()
	// fmt.Println(oddValue())
	// fmt.Println(oddValue())
	fmt.Println(fibo(-12))
}

func add(arr []float64)float64{
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	if len(arr) == 0 {
		panic("Invalid arr")
	}
	var sum float64
	for _,num:= range arr {
		sum += num
	}
	return sum

}

func half(num int)( int, bool){
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	var answer int;
	var value bool;
	if num <= 0{
		panic("number is not valid")
	}
	answer = num / 2;
	if num % 2 == 0{
		value = true
	}else{
		value = false
	}
	return answer, value
}

func highestNumber(args ...float64) float64{
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	if len(args) <= 0 {
		panic("Args in empty")
	}
	value := args[0]
	for _, num := range args {
		if num > value{
			value = num
		} 
	}
	return value
}

func makeOddGenerator()func()uint{
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 1
		return
	}
}

func fibo(num int) int {
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	if num < 0 {
		panic("Number is not valid")
	}

	if num == 0 {
		return 0
	}else if num == 1{
		return 1
	}
	return fibo(num-1) + fibo(num-2)
}