
2015/01/02 14:37:10 /home/lily/Work/go/src/exersice_Code/1.2.015/log.go:18: hello

2015/01/02 14:37:10 /home/lily/Work/go/src/exersice_Code/1.2.015/log.go:19: test
]string{
		"a":"aa",
		"b":"bb",
	}
	fmt.Println(m2)

	m1["a"] = "xx"
	m1["x"] = "xx"

	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	}else {
		fmt.Println("key not found")
	}

	delete(m1, "a")
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
    }else {
		fmt.Println("key not found")
	}

	fmt.Println("============")

	for k,v:=range m1{
		fmt.Printf("%s ==== %s\n", k, v)
	}

	fmt.Println(len(m1),len(m2))
}
