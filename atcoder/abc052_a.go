package atcoder

import "fmt"

func abc052a() {
	var lenA, widB, lenC, widD int
	fmt.Scan(&lenA, &widB, &lenC, &widD)

	rectangle1 := lenA * widB
	rectangle2 := lenC * widD

	if rectangle1 >= rectangle2 {
		fmt.Println(rectangle1)
	} else {
		fmt.Println(rectangle2)
	}
}
