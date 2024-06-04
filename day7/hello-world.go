package main

import "fmt"

func main() {

	//Day7 - June4, 2024
	//Functions
	fmt.Println("1+2=", plus(1,2))

	fmt.Println("1+2+3.14=", plusPlus(1,2,3.14))


	//Multiple return value functions
	a,b,c := vals()
	fmt.Println("received both values: ", a,b,c)

	x,_,z:= vals()
	fmt.Println("received only one value: ",x,z)


	//Variadic functions
	//function can be called with any number of arguments. fmt.Println() is a avariadic function
	fmt.Println("sum of multiple numbers: ", sums(1,2,3,4,5,6,7,8,9))



}

func sums(nums ...int) int{
	sum:=0
	for _,num := range nums {
		sum+=num
	}
	return sum
}


func plus(x int, y int) int {
	return x + y
}

func plusPlus(x,y int, z float32) int {
	return x +y + int(z)
}

func vals() (int, int, int) {
	return 3,5,9
}
