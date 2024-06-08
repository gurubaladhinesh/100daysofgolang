package main

import (
	"fmt"
	"math/big"
)

func main() {

	//Day11 -  June 8, 2024
	//Practice
	//Calculate factorial of big Int using iteration

	n:=big.NewInt(100)
	fmt.Println("factorial: ",fact(n))

}

func fact(n *big.Int) *big.Int {
	f := big.NewInt(1)
	var one = big.NewInt(1)
	for ;n.Cmp(one) ==1; n = n.Sub(n,one) {
		f = f.Mul(f, n);
	}
	return f
}
