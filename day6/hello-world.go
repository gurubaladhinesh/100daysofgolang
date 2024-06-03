package main

import "fmt"


func main() {

	//Day6 - June3, 2024
	//Range

	nums := []int{2,5,3,1}
	sum := 0
	for _,num:= range nums {
		sum += num
	}
	fmt.Println("sum:",sum)

	for i, num := range nums {
		fmt.Println("index: ", i, "num: ", num)
	}

	kvs := map[string]string {"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Println("key:",k," value:", v)
	}

	for k:= range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i,c)
	}



}
