package main

import (
	"fmt"
)

func main() {
/*	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(myArray)

	//var mySlice []int = myArray[:5]
	//fmt.Println(mySlice)

	//for s, _ := range mySlice {
		//fmt.Println(s, "")
	//}

	s := make([]string, 3)
	//fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//fmt.Println("set:", s)
	//fmt.Println("get:", s[2])
	//fmt.Println("len:", len(s))

	c := make([]string, len(s))
	//fmt.Println(c)
	copy(c, s)
	//fmt.Println(c)
	s = append(s, "d")
	//fmt.Println(s)
	s = append(s, "e", "f")
	//fmt.Println(s)

	l := s[2:5]
	//fmt.Println(l)
	l = s[:5]
	//fmt.Println(l)

	l = s[2:]
	//fmt.Println(l)

	t := []string{"g", "h", "i"}
	//fmt.Println(t)*/

	twoD := make([][]int, 3)
	//fmt.Println(twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		fmt.Println(innerLen)
		twoD[i]= make([]int, innerLen)
		fmt.Println(twoD[i])
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i+j
			 for _,v :=range twoD[i][j]{
				 fmt.Println("v:",v)
			 }
			fmt.Println(twoD[i][j])
		}
	}
	//fmt.Println("2d:", twoD)
}
