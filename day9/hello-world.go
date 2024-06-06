package main

import (
	"fmt"
	"math/big"
)

func main() {

	//Day9 - June6, 2024
	//Recursive Functions

	var fib func(n int)int
	
	fib = func (n int) int {
		if n<=2 {
			return n-1
		}
		return fib(n-1)+fib(n-2)
	}

	fmt.Println("fibonacci", fib(7))
	fmt.Println("factorial", fact(*big.NewInt(200)))

}


//TODO - revisit this after learning pointers
func fact(n big.Int) *big.Int{
	one := big.NewInt(1)
	cmp:= n.Cmp(one)/*n.Cmp(one)*/
	if cmp == -1 || cmp ==0 {
		return &n
	}
	sub:=n.Sub(&n,one)
	mul:= n.Mul(&n,fact(*sub))
	return mul
}



