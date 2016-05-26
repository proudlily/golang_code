package main

import (
	"fmt"

	"regexp"
)

func main() {
	a := "13260567913"
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]{9})$`, a); m {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
