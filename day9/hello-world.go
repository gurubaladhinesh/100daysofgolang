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
	fmt.Println("factorial", fact(big.NewInt(5)))

}

//TODO
func fact(n *big.Int) *big.Int{
	one := big.NewInt(1)
	cmp:= n.Cmp(one)/*n.Cmp(one)*/
	if cmp == -1 || cmp ==0 {
		return n
	}
	mul:= n.Mul(n,fact(n.Sub(n,one)))
	return mul
}



