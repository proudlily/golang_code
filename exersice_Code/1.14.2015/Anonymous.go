package main

import "fmt"

type Person struct {
	name string
	age  int
	contact struct {
		phone    string
		address  string
		postcode string
	}
}

func main() {
	d := Person{
		name:"lily",
		age :10,
	}

	d.contact.address="wuhan"
	d.contact.phone="13260567913"
	d.contact.postcode="00000"


	fmt.Println(d)
}
