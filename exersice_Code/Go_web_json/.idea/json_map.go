package main

import("fmt"
)

func main() {
	m := map[string]string{"key1":"val1"}
	fmt.Println(m)
	m["key2"] = "val2"
	fmt.Println(m)
	p := m["key1"]
	fmt.Println(p)
	for a,b := range m{
		fmt.Println("a=", a, "b=", b)
	}
	fmt.Println(m)
}
