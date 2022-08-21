package atcoder

import "fmt"

func abc055a() {
	var amountOfAte int
	fmt.Scan(&amountOfAte)

	payment := amountOfAte * 800
	cashBack := (amountOfAte / 15) * 200

	totalCost := payment - cashBack

	fmt.Println(totalCost)
}
