package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func(){
		for j := 0; j < 10; j++ {
			for i := 0; i < 10; i++ {
				c <- i
			}
		
	}
	}()
	for k := 0; k <100; k++ {
		fmt.Println(k, <-c)
	}
	fmt.Println("about to exit")
}
//Create a program that launches 10 goRoutines each adding 10 numbers to a channel
func lGo(){
	m := make(chan int)

	for i:= 0;i<10;i++{
		go func(){
			for j:=0;j<10;j++{
				m <- j
			}
		}() 
	}
	for k:=0;k<100;k++{
		fmt.Println(k,<-m)
	}
}
















//TODO: Exercise Number 1////
//FIXME:  Using a Func literal Make a channel run and then buffered it

func ej1() {
	x := make(chan int)
	go func() {
		x <- 2
		x <- 1
		x <- 0
		close(x)
	}()
	for v := range x {
		fmt.Println("Value:", v)
	}
}

//TODO: Exercise Number ////
//FIXME: Pull the Values of the Chanel with a Select Statement

// This function make a Channel and Fill it
func gen_4(q chan<- int) <-chan int {
	//Create Chanel
	c := make(chan int)
	//Call go func to make the channel Unblock
	go func() {
		//Fill Up the channel
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 0
		//Close the chanel so it is not listening
		close(c)
	}()
	//Return the Chanel
	return c
}

// This function take 2 channels then print them
func rec_4(q, c <-chan int) {
	for {
		select {
		case v := <-q:
			fmt.Println("The value is in the q Channel:", v)
		case <-c:
			return
		}
	}
}

//TODO: //	The Ok statement
//FIXME:// THe Correct use of the ",Ok" statement

func eaks() {
	c := make(chan int)
	//The channel need data inside it
	//You need the //GO\\ funcion to unblock the channels
	go func() {
		c <- 1
	}()
	//Inciialized the Ok statement
	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	fmt.Println("Now it's close")
	v, ok = <-c
	fmt.Println(v, ok)
}

// Write a Program that puts in 100 Number in a Channel and then print them
func r() {
	m := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			m <- i
		}
		close(m)
	}()

	for v := range m {
		fmt.Println("This is the value", v)
	}
}
