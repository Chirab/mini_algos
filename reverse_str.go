package main

import (
	"strings"
	"fmt"
)

func reverseStr(s string) string {
	var res strings.Builder
	for i:= len(s) - 1; i > -1; i-- {
		res.WriteString(string(s[i]))
	}

	return res.String()
}

func main() {

	input := "Hello World"
	result := reverseStr(input)
	fmt.Println(result)
}
