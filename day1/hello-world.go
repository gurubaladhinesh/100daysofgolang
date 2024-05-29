package main

import (
	"fmt"
	"math"
)

const pi float32=3.14

func main() {
	//hello-world
	fmt.Println("Hello World")

	//values
	fmt.Println("abc "+"def")
	fmt.Println("1+2=",1+2)
	fmt.Println("7.0/3.0=",7.0/3.0)
	fmt.Println(true || false)
	fmt.Println(true && false)
	fmt.Println(!true)


	//variables
	var a = "initial"
	var b,c int = 9, 56
	var d int
	var e string
	var f float32=4.23
	var g bool = true
	var h bool
	i:= 2.11
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	//constants
	const radius float32=4
	fmt.Println(radius)
	fmt.Println("area:", pi * radius * radius)
	fmt.Println(int64(12))
	fmt.Println(fmt.Sprint(123243))
	fmt.Println(string("123243"))
	fmt.Println(math.Sin(float64(radius)))

}