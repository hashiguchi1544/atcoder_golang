package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	fmt.Println(strings.ReplaceAll(input, ",", " "))

}
