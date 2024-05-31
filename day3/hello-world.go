package main

import "fmt"


func main() {
	
	//continue with switch

	whatAmI := func(i interface{}) {
		switch tp:= i.(type) { //type keyword cannot be used with if , it can be only used with switch
		case bool:
			fmt.Println("it is Boolean")
		case int:
			fmt.Println("it is Int")
		default:
			fmt.Printf("unknown type %T\n",tp)
		}
	}
	whatAmI(true)
	whatAmI(10)
	whatAmI("abc")


	//Arrays
	var a [5]int;
	fmt.Println("emp: ", a)
	fmt.Println("emp len: ", len(a))
	a[3]=13
	fmt.Println("after setting: ", a)

	b:=[5]int{1,2,3,4,5}
	fmt.Println("arr b: ", b)

	c:=[...]int{2,3,4,5,6,7}
	fmt.Println("arr c len: ", len(c))

	d:=[...]int{100,3:400,500} //: colon mentions to put 400 on index 3
	fmt.Println("arr d: ",d)

	var twoD [2][3]int
	for i:=0;i<2;i++ {
		for j:=0;j<3;j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("twoD: ",twoD)

	twoDD:=[...][3]int{ //... notation can only resolve the length of array upto 1 level. cannot do that to next level
		{1,2,3},
		{4,5,6},
	}
	fmt.Println("twoDD: ", twoDD)


}