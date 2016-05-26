package main

import "fmt"

func args()(int,string){
	return 1,"abc"
}
func test(i int,s string) {
   fmt.Println(i,s)
}
func main(){
	test(args())
}
