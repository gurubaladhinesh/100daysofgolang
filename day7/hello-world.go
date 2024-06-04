package main

import "fmt"

func main() {

	//Day7 - June4, 2024
	//Functions
	fmt.Println("1+2=", plus(1,2))

	fmt.Println("1+2+3.14=", plusPlus(1,2,3.14))


	//Multiple return value functions

}


func plus(x int, y int) int {
	return x + y
}

func plusPlus(x,y int, z float32) int {
	return x +y + int(z)
}
