package main

import (
	"math/rand"
	"fmt"
	"time"
)

//go concurrency

// to make a function run concurrently add the go keyword before the function call
func main(){

	// select is a switch case but only works for channels
	c1 := make(chan int,20)
	c2 := make(chan int,  5)

	go func(){
		for {
			c1 <- 50
			time.Sleep(time.Second * 2)
		}
	}()

	go func(){
		for {
			c2 <- 80
			time.Sleep(time.Second * 3)
		}
	}()

	go func(){
		for {
			select{
			case msg1 := <- c1:
				fmt.Println("message 1",msg1)
			case msg2 := <- c2:
				fmt.Println("message 2",msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	

	var input string
	fmt.Scanln(&input)
}
	


//go channel 
// this is a way of making two gorrountines talk to each other key word (chan)
// you can make the function only to send or recieve
// c chan <- send and c <- chan recieve
// a channel with out direction is bi-dirrectional
func f(n int){
	num:=[]int{1,2,3,4,5,6,7,8,9,10}
	for _, i:= range num{
		fmt.Println(n, ": ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger (c chan <- string){
	for i:=0; ; i++{
		c <- "ping"
	}
}

func printer(c <- chan string){
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep((time.Second * 1))
	}
}

func ponger(c chan <- string){
	for i:=0; ; i++ {
		c <- "pong"
	}
}