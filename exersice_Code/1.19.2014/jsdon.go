package main

import (
	"bytes"
	"encoding/json"

	"fmt"

)

func main() {
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	decoder := json.NewDecoder(buffer)

	data := map[string]interface{}{
		"s":"Hello,World!",
		"i":1234,
		"m":map[string]int{"x":1, "y":2},
	}
	encoder.Encode(data)

	fmt.Printf("%v\n", buffer.String())

	d := make(map[string]interface{})
	decoder.Decode(&d)// 注意使用用指针,因为转换成 interface{} 时发生生 value-copy。
	fmt.Printf("%#v\n", d)
}
