package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

//learning about the inbuilt packages in go
//string
//fmt
//os
//io/ioutil is already deprecated use io or os

func main() {
	fmt.Println(strings.Compare("test","test2"), strings.ToUpper("walon"))
	// file,err := os.Open("test.txt")
	// if err !=nil{
	// 	return
	// }
	// defer file.Close()

	// //getting file size
	// stat,err := file.Stat()
	// if err != nil{
	// 	return
	// }

	// //reading the file
	// bs := make([]byte, stat.Size())
	// _,err = file.Read(bs)
	// if err != nil{
	// 	return
	// }
	// str:= string(bs)
	// fmt.Println(str)

	bs,err := ioutil.ReadFile("test.txt")
	if err !=nil{
		return
	}
	str := string(bs)
	fmt.Println(str)

	file,err := os.Create("test2.txt")
	if err !=nil{
		return
	}
	defer file.Close()
	file.WriteString("hello crazyyyy")

}