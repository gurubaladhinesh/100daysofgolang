package main

import (
	"fmt"
	"time"
)

func main() {
	
	//for loop
	i:=1
	fmt.Println("loop1")
	for i<=3 {
		fmt.Println(i)
		i++;
	}

	fmt.Println("loop2")
	for j:=0; j<=3;j++ {
		fmt.Println(j)
	}

	fmt.Println("loop3 range i:", i)
	for i:= range 3 {
		fmt.Println("range",i)
	}

	fmt.Println("loop4")
	for {
		fmt.Println("infinite loop")
		break
	}

	fmt.Println("loop5")

	for n:= range 7 {
		if n%2==0 {
			continue
		}
		fmt.Println(n)
	}

	//if-else

	if 7%2==0 {
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if n:=9; n<0 {
		fmt.Println("n is negative")
	}

	if m,n:= 1,2; m<n {
		fmt.Println("m < n")
	}else if( m == n){
		fmt.Println("m == n")
	}else{
		fmt.Println("m > n")
	}

	//switch
	switch num:=3;num { //same notation used as in line#54 in if statemet. same notation can be used in for loop
	case 1,3:
		fmt.Println("one,three")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("number")
	}

	tim:= time.Now()

	switch{
	case tim.Hour() < 12:
		fmt.Println("morning")
	case tim.Hour() == 12:
		fmt.Println("noon")
	case tim.Hour()>12 && tim.Hour() < 18:
		fmt.Println("evening")
	case tim.Hour() > 18:
		fmt.Println("night")
	default:
		fmt.Println("unknown time")
	}

}