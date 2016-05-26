package main

import (
	"fmt"

	"regexp"
	"strings"
)

func main() {
	a := "ASCDDFDSCGDGGGGSSS"

	/*resp, err := http.Get(a)
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}
	src := string(body)*/

	re, _ := regexp.Compile(`D{1,}`)
	src := re.ReplaceAllString(a, "\n")
	fmt.Println(strings.TrimSpace(src))

}
