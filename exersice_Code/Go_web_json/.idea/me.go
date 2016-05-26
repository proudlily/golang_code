package main

import (
	"fmt"

)
func a (b *int){
	*b = *b - 1
}
func main(){
	b :=10
	 a(&b)
	fmt.Println(b)
}
