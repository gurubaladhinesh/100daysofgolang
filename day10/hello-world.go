package main

import (
	"fmt"
)

func main() {

	//Day10 - June7, 2024
	//Pointers
	n :=1
	incVal(n)
	fmt.Println("incVal", n)
	incRef(&n)
	fmt.Println("incRef",n)

}

func incVal(n int) {
	n++
}
func incRef(n *int) {
	*n++;
}



