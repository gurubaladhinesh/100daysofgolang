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

	fmt.Println("parameter passing closure")
	pos, neg:=adder(),adder()
	for i:=0;i<10;i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
	
}

func initSeq() func() int {
	i:=0
	return func ()  int {
		i++;
		return i;
	}
}

func adder() func(int) int {
	sum:=0
	return func(i int) int {
		sum+=i
		return sum
	}
}
