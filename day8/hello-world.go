package main

import "fmt"

func main() {

	//Day8 - June5, 2024
	//Closures

	nextInt:=initSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts:= initSeq()
	fmt.Println(newInts())
	
}

func initSeq() func() int {
	i:=0
	return func ()  int {
		i++;
		return i;
	}
}
