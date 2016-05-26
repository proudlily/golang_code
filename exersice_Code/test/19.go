package main

func test(x *[4]int) {
	for i := 0; i < len(x); i++ {
		println(x[i])
	}
	x[3] = 300
}

func main() {
	x := &[4]int{2: 100, 1: 200}
	test(x)
	println(x[3])
}
