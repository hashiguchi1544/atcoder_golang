package atcoder

import "fmt"

func abc053a() {
	var score int
	fmt.Scan(&score)

	if score < 1200 {
		fmt.Println("ABC")
	} else {
		fmt.Println("ARC")
	}

}
