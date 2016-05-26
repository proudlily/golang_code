package main

func main() {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				break LABEL1
			} else {
				println("L1:", i)
			}
		}
	}
LABEL2:
	for i := 0; i < 5; i++ {
		for {
			println("L2:", i)
			continue LABEL2
		}
	}
	println("over")
}
