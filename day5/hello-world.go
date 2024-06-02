package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {

	//slices - example

	var s []string
	fmt.Println("unitt::", s, s == nil, len(s) == 0, len(s), cap(s))

	fmt.Printf("%T\n", s)

	s = make([]string, 3)
	fmt.Println("unitt slice::", s, "len:", len(s), "cap:", cap(s))

	s[0]="a"
	s[1]="b"
	s[2]="c"

	fmt.Println("set:",s)
	fmt.Println("get:",s[1])
	s = append(s, "d")
	s= append(s, "e","f", "g")
	fmt.Println("appended s:",s, len(s), cap(s))

	c:=make([]string, len(s))
	copy(c,s)
	fmt.Println("copied slice c: ",c, len(c), cap(c))

	l := s[2:5]
	fmt.Println("sliced l: ", l)

	l = s[:3]
	fmt.Println("l: ", l)

	l = s[2:]
	fmt.Println("l: ", l)

	t:=[]string{"z","x","c"}
	fmt.Println("t: ", t)

	t2:=[]string{"z","x","c"}
	if slices.Equal(t,t2) {
		fmt.Println("t == t2")
	}

	var matrix = make([][]int, 5)
	for i:=0;i<5;i++ {
		matrix[i] = make([]int, i+1)
		for j:=0;j<=i;j++ {
			matrix[i][j] = i+j
		}
	}
	fmt.Println("matrix:: ", matrix)


	//Maps

	var nmap = make(map[string]int)
	nmap["a1"]=1
	nmap["a2"]=2
	nmap["a3"]=3
	fmt.Println("nmap: ", nmap)

	v1:=nmap["a1"]
	fmt.Println("v1: ", v1)

	v4 := nmap["a4"]
	fmt.Println("v4: ", v4)// if key doesn't exist 0(zero) is returned

	fmt.Println("map len: ", len(nmap))

	delete(nmap, "a2")
	fmt.Println("deleted: ", nmap)


	_, prs := nmap["a2"]
    fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }

}
