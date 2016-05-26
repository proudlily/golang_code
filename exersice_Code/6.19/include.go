package main

import "fmt"

type Bird struct {
	fly     string
	feather string

	sing string
}
type Duck struct {
	Bird
	taohaozhuren string
}

func (h *Bird) Play() {
	fmt.Printf("Hi,My way is %s!", h.fly)
	fmt.Printf("I have  %s feather!", h.feather)
	fmt.Printf("I can %s!", h.sing)

}

func main() {
	b := Duck{Bird{"fly", "yello", "gagaga"}, "Golang Inc"}

	b.Play()

}
