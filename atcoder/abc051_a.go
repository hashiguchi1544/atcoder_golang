package atcoder

import (
	"fmt"
	"strings"
)

func abc051a() {
	var input string
	fmt.Scan(&input)

	fmt.Println(strings.ReplaceAll(input, ",", " "))

}
