package main

import "fmt"

func main() {
	var score int
	fmt.Scan(&score)

	if score < 1200 {
		fmt.Println("ABC")
	} else {
		fmt.Println("ARC")
	}

}